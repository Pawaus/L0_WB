# L0 task
### Env file:
- POSTGRES_USER - Initial user in database(default: postgres)
- POSTGRES_PASSWORD - Initial password for user in database(default: postgres) 
- POSTGRES_DB - Initial name database(default: postgres)
- APP_DB_HOST - Container name database(from docker-compose -> container_name of database service, default: db)
- APP_DB_PORT - Database port (default: 5432)
- APP_DB_USER - Database user (default: postgres)
- APP_DB_PASSWORD - Database password (default: postgres)
- APP_DB_NAME - Database name (dafault: postgres)
- NATS_CLUSTER_NAME - Nats-streaming cluster name(default: test-cluster)
- NATS_CLIENT_NAME - Nats-streaming client name (default: client-123)
- NATS_SERVER_NATS - Nats-streaming server(from docker-compose -> container_name of nats-streaming service, default: nats)
### Run app:
```docker compose up --build```
### Push message to nats-streaming channel
```go run ./cmd/push```