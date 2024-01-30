package app

import (
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/configuration"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Config *configuration.Configuration
}

var appInstance *App

func init() {
	appInstance = &App{}
}

func Get() *App {
	return appInstance
}
