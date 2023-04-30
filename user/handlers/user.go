package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"user-api/exception"
	request2 "user-api/model/request"
	"user-api/service"
)

func GetUser(request events.APIGatewayV2HTTPRequest, service service.UserService) (*events.APIGatewayV2HTTPResponse, error) {
	userId := request.PathParameters["userId"]
	if len(userId) == 0 {
		return nil, exception.InvalidInputError
	}
	user, err := service.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return apiResponse(http.StatusOK, user)
}

func DeleteUser(request events.APIGatewayV2HTTPRequest, service service.UserService) (*events.APIGatewayV2HTTPResponse, error) {
	userId := request.PathParameters["userId"]
	if len(userId) == 0 {
		return nil, exception.InvalidInputError
	}
	err := service.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	return apiResponse(http.StatusOK, nil)
}

func CreateUser(request events.APIGatewayV2HTTPRequest, userService service.UserService) (*events.APIGatewayV2HTTPResponse, error) {
	var userRequest request2.UserCreateRequest
	if err := json.Unmarshal([]byte(request.Body), &userRequest); err != nil {
		return nil, exception.InvalidInputError
	}

	err := userService.SaveUser(userRequest)
	if err != nil {
		return apiResponse(500, nil)
	}
	return apiResponse(http.StatusCreated, nil)
}

func apiResponse(status int, body interface{}) (*events.APIGatewayV2HTTPResponse, error) {
	resp := events.APIGatewayV2HTTPResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	if body != nil {
		stringBody, _ := json.Marshal(body)
		resp.Body = string(stringBody)
	}
	return &resp, nil
}
