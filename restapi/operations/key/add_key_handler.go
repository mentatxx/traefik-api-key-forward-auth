package key

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/app"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
)

func AddKeyHandlerImpl(params AddKeyParams, principal *models.AuthPrincipal) middleware.Responder {
	db := app.Get().DB

	key := models.Key{
		Key:            uuid.NewString(),
		OrganizationID: params.Body.OrganizationID,
		ProjectID:      params.Body.ProjectID,
		UserID:         params.Body.UserID,
		Attributes:     params.Body.Attributes,
	}
	result := db.Create(&key)
	if result.Error != nil {
		return middleware.Error(500, result.Error)
	}
	return NewAddKeyOK().WithPayload(&models.CreateKeyResult{
		ID:  key.ID,
		Key: key.Key,
	})
}
