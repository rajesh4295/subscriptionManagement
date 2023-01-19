package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"subscriptionManagement/model"
	"subscriptionManagement/route/v1"
	"subscriptionManagement/util"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	serverRunMode := "dev"
	if mode, ok := os.LookupEnv("ENV"); ok {
		serverRunMode = mode
	}

	fmt.Printf("Server mode: %s", serverRunMode)
	fmt.Println()
	appCtx, err := initServer(serverRunMode)

	if err != nil {
		fmt.Printf("Server init failed: %s", err)
		fmt.Println()
		os.Exit(1)
	}
	baseRoute := appCtx.Router.Group("/api")
	v1Route := baseRoute.Group("/v1")
	{
		route.InitHealthRoute(v1Route)
		route.InitUserRoute(v1Route)
		route.InitProductRoute(v1Route, appCtx)
	}
	err = appCtx.DB.AutoMigrate(
		&model.Plan{},
		&model.Product{},
		&model.Subscription{},
		&model.Tenant{},
		&model.Membership{},
		&model.User{},
	)
	if err != nil {
		appCtx.Logger.Error(err)
		os.Exit(1)
	}

	appCtx.Logger.Info("Server init complete")

	port := appCtx.Env.Get("app.port")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: appCtx.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appCtx.Logger.Errorf("Error starting server %s", err)
		}
	}()

	appCtx.Logger.Infof("Server is running on port %d", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	appCtx.Logger.Warn("Shutdown init")

	db, _ := appCtx.DB.DB()
	db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appCtx.Logger.Fatal("Shutdown failed:", err)
	}

	appCtx.Logger.Warn("Shutdown complete")
}

func initServer(mode string) (*model.AppCtx, error) {
	gin.SetMode(gin.ReleaseMode)
	appCtx := &model.AppCtx{
		Router: gin.New(),
	}
	// Init Env provider
	env, err := util.InitEnv(mode)
	if err != nil {
		return nil, err
	}
	appCtx.Env = env

	// Init log provider
	logger, err := util.InitLogger(env.Get("logLevel").(string))
	if err != nil {
		return nil, err
	}
	baseLogger := logger.WithField("mode", mode)
	appCtx.Logger = baseLogger.Logger

	// Init Database
	db, err := util.InitDB(env)
	if err != nil {
		return nil, err
	}
	appCtx.DB = db

	return appCtx, nil
}
