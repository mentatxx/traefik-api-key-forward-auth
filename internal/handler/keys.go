package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/repository"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
	"github.com/mentatxx/traefik-api-key-forward-auth/restapi/operations/key"
)

type KeysHandlerStruct struct {
	KeyAddKeyHandler    key.AddKeyHandlerFunc
	KeyDeleteKeyHandler key.DeleteKeyHandlerFunc
	KeyGetKeysHandler   key.GetKeysHandlerFunc
}

func NewKeysHandler(repo repository.KeysRepository) *KeysHandlerStruct {
	return &KeysHandlerStruct{
		KeyGetKeysHandler: func(params key.GetKeysParams, principal *models.AuthPrincipal) middleware.Responder {
			return getKeysHandlerImpl(repo, params, principal)
		},
		KeyAddKeyHandler: func(params key.AddKeyParams, principal *models.AuthPrincipal) middleware.Responder {
			return AddKeyHandlerImpl(repo, params, principal)
		},
		KeyDeleteKeyHandler: func(params key.DeleteKeyParams, principal *models.AuthPrincipal) middleware.Responder {
			return DeleteKeyHandlerImpl(repo, params, principal)
		},
	}
}
