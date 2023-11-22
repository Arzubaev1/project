
run:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migration/postgres -database 'postgres://postgres:abdu04abdu@localhost:5432/taxi_sharing?sslmode=disable' up

migration-down:
	migrate -path ./migration/postgres -database 'postgres://postgres:abdu04abdu@localhost:5432/taxi_sharing?sslmode=disable' down


