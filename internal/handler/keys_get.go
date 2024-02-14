package handler

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/repository"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
	"github.com/mentatxx/traefik-api-key-forward-auth/restapi/operations/key"
)

func getKeysHandlerImpl(repo repository.KeysRepository, params key.GetKeysParams, principal *models.AuthPrincipal) middleware.Responder {
	var filter repository.KeyFilter
	err := json.Unmarshal([]byte(params.Filter), &filter)
	if err != nil {
		return middleware.Error(400, "Invalid filter")
	}

	keys, err := repo.GetByFilter(filter, params.After, params.Limit)

	if err != nil {
		return middleware.Error(500, err)
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
	return key.NewGetKeysOK().WithPayload(resultKeys)
}
