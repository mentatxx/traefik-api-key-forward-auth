package key

import (
	"encoding/json"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/app"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
)

type KeyFilter struct {
	OrganizationID *string `json:"organizationId,omitempty"`
	ProjectID      *string `json:"projectId,omitempty"`
	UserID         *string `json:"userId,omitempty"`
}

func GetKeysHandlerImpl(params GetKeysParams, principal *models.AuthPrincipal) middleware.Responder {
	db := app.Get().DB
	var keys []models.Key
	var filter KeyFilter
	err := json.Unmarshal([]byte(params.Filter), &filter)
	if err != nil {
		return middleware.Error(400, err)
	}

	result := db.Where(filter).Find(&keys)
	if result.Error != nil {
		return middleware.Error(500, result.Error)
	}
	return NewGetKeysOK().WithPayload(&models.GetKeysResult{})
}
