package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/repository"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
	"github.com/mentatxx/traefik-api-key-forward-auth/restapi/operations/key"
)

func DeleteKeyHandlerImpl(repo repository.KeysRepository, params key.DeleteKeyParams, principal *models.AuthPrincipal) middleware.Responder {
	error := repo.MarkAsDeletedByID(params.ID)
	if error != nil {
		return middleware.Error(500, error)
	}
	return key.NewDeleteKeyOK()
}
