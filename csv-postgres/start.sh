docker run -e DB_USER=keycloak \
	-e DB_PASSWORD=password \
	-e DB_NAME=keycloak \
	-e DB_HOST=postgresql \
	-e DB_PORT=5432 \
	csv-to-postgres
