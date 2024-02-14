generate-server:
	swagger generate server -f ./swagger/swagger.yml -A traefik-api-key-forward-auth --principal models.AuthPrincipal

build:
	go build -o ./dist/traefik-api-key-forward-auth-server ./cmd/traefik-api-key-forward-auth-server
