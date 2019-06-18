package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/edds/businesslogic/domain"
)

type DomainAdapter interface {
	CalculateDeceptiveScore(request events.APIGatewayProxyRequest) (Response, error)
}

type domainAdapter struct {
	domainService domain.PrimaryPort
}

func NewDomainAdapter(domainService domain.PrimaryPort) DomainAdapter {
	return &domainAdapter{
		domainService,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func (a *domainAdapter) CalculateDeceptiveScore(request events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	//Get the path parameter that was sent
	domainname := request.PathParameters["domain"]
	if domainname == "" {
		return Response{StatusCode: 400}, nil
	}

	d := domain.Domain{Name: domainname}

	score, err := a.domainService.CalculateDeceptiveScore(&d)
	fmt.Println("score", score)

	if err != nil {
		return Response{StatusCode: 502}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"score": score,
	})

	if err != nil {
		return Response{StatusCode: 400}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "GET",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}

	return resp, nil
}
