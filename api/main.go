package main

import (
	"github.com/FernandoCagale/serverless-auth/api/routers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/negroni"
	"github.com/urfave/negroni"
)

var initialized = false

var negroniLambda *negroniadapter.NegroniAdapter

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !initialized {
		router := routers.NewRouter()

		routers.MakeHandlers(router)

		n := negroni.Classic()
		n.UseHandler(router)

		negroniLambda = negroniadapter.New(n)
		initialized = true
	}

	return negroniLambda.Proxy(req)
}

func main() {
	lambda.Start(handlers)
}
