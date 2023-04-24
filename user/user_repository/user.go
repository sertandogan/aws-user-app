package user_repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"user-api/domain"
	"user-api/exception"
)

const AwsTable = "user"

type userRepository struct {
	awsClient *dynamodb.DynamoDB
}

type UserRepository interface {
	AddUser(request *domain.UserEntity) error
	GetUser(userId string) (*domain.UserEntity, error)
}

func NewUserRepository(client *dynamodb.DynamoDB) UserRepository {
	return &userRepository{awsClient: client}
}

func (repo *userRepository) AddUser(entity *domain.UserEntity) error {
	attributeValue, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return exception.InvalidInputError
	}
	userEntityItem := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(AwsTable),
	}

	_, err = repo.awsClient.PutItem(userEntityItem)
	if err != nil {
		log.Default().Printf("Got error calling PutItem: %s", err)
		return exception.UnknownException
	}
	return nil
}

func (repo *userRepository) GetUser(userId string) (*domain.UserEntity, error) {
	result, err := repo.awsClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(AwsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
	})

	if err != nil {
		log.Default().Printf("Got error calling GetUser: %s", err)
		return nil, exception.UnknownException
	}
	if result.Item == nil {
		return nil, exception.UserNotFound
	}

	user := new(domain.UserEntity)
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, exception.InvalidInputError
	}
	return user, nil
}
