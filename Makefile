run: *.go swag
	go run main.go

deps:
	go get -u github.com/swaggo/swag/cmd/swag

swag: */*.go
	swag init


