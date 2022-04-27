## Service provide an API for monitoring query speed
API specification can be found in api/openapi-spec folder

### Before executing service make sure to add below changes to postgres config file

- shared_preload_libraries = 'pg_stat_statements'
- max_connections=200
- pg_stat_statements.track = all

### OR just run docker-compose file

## Create your .env file with below variables
 - DB_HOST=postgres
 - DB_NAME=postgres
 - DB_PASSWORD=12345
 - DB_PORT=5432
 - DB_USER=postgres
 - SERVICE_PORT=8080

## run service with docker-compose for not download any dependency
docker-compose --env-file .env up -d

service will be run on 127.0.0.1:8080

