ritter:
	go build -race -o ./bin/ritter ./cmd/ritter

run:
	go run pafaul/ritter

gen-db:
	sqlc generate

migrate-up: gen-db
	migrate -source file://sqlc/migrations -database sqlite3://db.sqlite up

migrate-down:
	migrate -source file://sqlc/migrations -database sqlite3://db.sqlite down

gen-swagger:
	swag init -g cmd/ritter/main.go