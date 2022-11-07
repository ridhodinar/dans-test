generate:
	swagger generate server -f api/swagger.yaml --default-scheme http

run:
	make generate
	go run cmd/dans-server/main.go --port=8080