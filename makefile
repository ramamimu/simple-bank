init:
	docker compose -f ./docker-compose/postgres.yaml up -d
	docker exec -ti postgres_simplebank psql -U postgres -c "CREATE DATABASE simplebank"
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose up

startdb:
	docker compose -f ./docker-compose/postgres.yaml up -d

stopdb:
	docker compose -f ./docker-compose/postgres.yaml down

createdb:
	docker exec -ti postgres_simplebank psql -U postgres -c "CREATE DATABASE simplebank"

dropdb:
	docker exec -ti postgres_simplebank psql -U postgres -c "DROP DATABASE simplebank"

migrateup:
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose up

migratedown:
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose down

migrateforce:
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose force $(version)

sqlc:
	sqlc generate