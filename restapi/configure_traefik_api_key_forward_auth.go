// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/mentatxx/traefik-api-key-forward-auth/internal/app"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/configuration"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/database"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
	"github.com/mentatxx/traefik-api-key-forward-auth/restapi/operations"
	"github.com/mentatxx/traefik-api-key-forward-auth/restapi/operations/key"
)

//go:generate swagger generate server --target ../../traefik-api-key-forward-auth --name TraefikAPIKeyForwardAuth --spec ../swagger/swagger.yml --principal AuthPrincipal

func configureFlags(api *operations.TraefikAPIKeyForwardAuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TraefikAPIKeyForwardAuthAPI) http.Handler {
	configureApp()

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-Api-Key" header is set
	api.ManagementAuthAuth = func(token string) (*models.AuthPrincipal, error) {
		config := app.Get().Config
		if config.ApiSecret == token {
			return &models.AuthPrincipal{
				Authenticated: true,
			}, nil
		}
		return nil, errors.Unauthenticated("http")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.KeyAddKeyHandler = key.AddKeyHandlerFunc(key.AddKeyHandlerImpl)
	api.KeyDeleteKeyHandler = key.DeleteKeyHandlerFunc(key.DeleteKeyHandlerImpl)
	api.KeyGetKeysHandler = key.GetKeysHandlerFunc(key.GetKeysHandlerImpl)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func configureApp() {
	// Initialize the app here
	app := app.Get()
	config := configuration.New()
	app.Config = config
	gormDb, err := database.Connect(config)
	if err != nil {
		panic(err)
	}
	app.DB = gormDb

}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
