clean:
	rm -rf ./bin

build-all: build-notes-api

build-notes-api:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/notes-api ./cmd/notes-api/lambda/main.go

deploy-dev-notes-api:
	SLS_DEBUG=* sls deploy function -f notes-api -s dev --force

deploy-dev-all:
	SLS_DEBUG=* sls deploy -s dev --force

deploy-prod-notes-api:
	SLS_DEBUG=* sls deploy function -f notes-api -s prod --force

deploy-prod-all:
	SLS_DEBUG=* sls deploy -s prod --force
