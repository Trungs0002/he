include .env
export

## Run the API
run-api:
	go run cmd/api/main.go

test:
	go test -mod=vendor -coverprofile=c.out -failfast -timeout 5m ./...

swagger:
	swag init -g cmd/api/main.go -o docs
