package key

import (
	"encoding/json"
	"log"

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
		return middleware.Error(400, "Invalid filter")
	}

	chain := db.Where(filter).Order("id")
	if (params.After != nil) && (*params.After != "") {
		chain = chain.Where("id > ?", *params.After)
	}
	if (params.Limit != nil) && (*params.Limit > 0) {
		chain = chain.Limit(int(*params.Limit))
	} else {
		chain = chain.Limit(50)
	}

	result := chain.Find(&keys)
	if result.Error != nil {
		return middleware.Error(500, result.Error)
	}
	// map result
	var resultKeys []*models.GetKeysResult
	for _, key := range keys {
		attrs, ok := key.Attributes.(*interface{})
		if !ok {
			log.Println("GetKeysHandlerImpl: Can not cast Attributes")
			return middleware.Error(500, "Invalid attributes")
		}
		attrsBytesArray, ok := (*attrs).([]byte)
		if !ok {
			return middleware.Error(500, "Invalid attributes")
		}
		attrsObject := make(map[string]interface{})
		err := json.Unmarshal(attrsBytesArray, &attrsObject)
		if err != nil {
			return middleware.Error(500, err)
		}

		resultKeys = append(resultKeys, &models.GetKeysResult{
			ID:             key.ID,
			Key:            key.Key,
			OrganizationID: key.OrganizationID,
			ProjectID:      key.ProjectID,
			UserID:         key.UserID,
			Attributes:     attrsObject,
		})
	}
	return NewGetKeysOK().WithPayload(resultKeys)
}
