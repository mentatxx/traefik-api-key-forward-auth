generate-server:
	swagger generate server -f ./swagger/swagger.yml -A traefik-api-key-forward-auth --principal models.AuthPrincipal
