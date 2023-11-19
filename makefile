init:
	docker compose -f ./docker-compose/postgres.yaml up -d
	docker exec -ti postgres_simplebank psql -U postgres -c "CREATE DATABASE simplebank"
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose up

dbstart:
	docker compose -f ./docker-compose/postgres.yaml up -d

dbstop:
	docker compose -f ./docker-compose/postgres.yaml down

dbcreate:
	docker exec -ti postgres_simplebank psql -U postgres -c "CREATE DATABASE simplebank"

dbdrop:
	docker exec -ti postgres_simplebank psql -U postgres -c "DROP DATABASE simplebank"

migrateup:
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose up

migratedown:
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose down

migrateforce:
	migrate -path=./db/migration -database postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable -verbose force $(version)

test:
	go clean -testcache
	go test -v -cover ./...

sqlc:
	sqlc generate