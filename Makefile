clean:
	rm -rf ./bin

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/primes-api ./cmd/primes-api/lambda/main.go

deploy:
	SLS_DEBUG=* sls deploy function -f primes-api --force
