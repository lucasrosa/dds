package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lucasrosa/edds/adapters/database/memory"
	"github.com/lucasrosa/edds/businesslogic/domain"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	domainRepo := domainadaptermemory.NewMemoryDomainRepository()
	domainService := domain.NewDomainService(domainRepo)
	domainAdapter := NewDomainAdapter(domainService)

	return domainAdapter.CalculateDeceptiveScore(request)
}

func main() {
	lambda.Start(Handler)
}
