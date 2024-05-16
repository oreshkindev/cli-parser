migrate create -ext sql -dir database/migrations 'table_name'

migrate -path database/migrations -database 'postgres://postgres:postgres@localhost:5432/postgres' up

migrate -path database/migrations -database 'postgres://postgres:postgres@localhost:5432/postgres' down
