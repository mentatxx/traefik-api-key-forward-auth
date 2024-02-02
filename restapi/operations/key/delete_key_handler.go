package key

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
)

func DeleteKeyHandlerImpl(params DeleteKeyParams, principal *models.AuthPrincipal) middleware.Responder {
	return middleware.NotImplemented("operation key.DeleteKey has not yet been implemented")
}
