package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/repository"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
	"github.com/mentatxx/traefik-api-key-forward-auth/restapi/operations/key"
)

func AddKeyHandlerImpl(repo repository.KeysRepository, params key.AddKeyParams, principal *models.AuthPrincipal) middleware.Responder {

	newKey := models.Key{
		Key:            uuid.NewString(),
		OrganizationID: params.Body.OrganizationID,
		ProjectID:      params.Body.ProjectID,
		UserID:         params.Body.UserID,
		Attributes:     params.Body.Attributes,
	}
	error := repo.Create(&newKey)
	if error != nil {
		return middleware.Error(500, error)
	}
	return key.NewAddKeyOK().WithPayload(&models.CreateKeyResult{
		ID:  newKey.ID,
		Key: newKey.Key,
	})
}
