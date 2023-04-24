package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"user-api/handlers"
	"user-api/service"
	"user-api/user_repository"
)

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	fmt.Printf("request: %+v", request)
	switch request.RequestContext.HTTP.Method {
	case "POST":
		{
			return handlers.CreateUser(request, userService)
		}
	case "GET":
		{
			return handlers.GetUser(request, userService)
		}
	default:
		{
			return &events.APIGatewayV2HTTPResponse{
				Body:       fmt.Sprintf("Unsupported method type. %s", request.RouteKey),
				StatusCode: 405,
			}, nil
		}
	}

}

var (
	userService service.UserService
)

func main() {
	region := os.Getenv("AWSREGION")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: &region,
		},
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	userRepository := user_repository.NewUserRepository(svc)
	userService = service.NewUserService(userRepository)
	lambda.Start(handler)
}
