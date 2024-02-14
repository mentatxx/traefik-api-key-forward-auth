package app

import (
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/configuration"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/database"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/repository"
)

type App struct {
	DB             *database.Database
	Config         *configuration.Configuration
	KeysRepository repository.KeysRepository
}

var appInstance *App

func init() {
	appInstance = &App{}
}

func Get() *App {
	return appInstance
}
