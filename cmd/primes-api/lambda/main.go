package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/therealgaryj/go-lambda-utils/pkg/lambdatohttp"
	"github.com/therealgaryj/go-prime-api/internal/app/primes"
)

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	r := primes.Initialize()

	return lambdatohttp.ServeRequest(r, ctx, req), nil
}

func main() {
	lambda.Start(HandleRequest)
}