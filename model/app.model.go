package model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppCtx struct {
	DB     *gorm.DB
	Router *gin.Engine
	Env    *viper.Viper
	Logger *logrus.Logger
}

func (app *AppCtx) GetConfig() *AppCtx {
	return app
}
