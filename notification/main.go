package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
)

func handler(ctx context.Context, event events.DynamoDBEvent) error {
	for _, record := range event.Records {
		eventName := record.EventName

		switch eventName {
		case "INSERT":
			user, err := convertEventToUser(record.Change.NewImage)
			if err != nil {
				return err
			}
			err = sendEmail(user.Email, fmt.Sprintf("Welcome %s", user.Name), "Welcome Mail")
			if err != nil {
				return err
			}
		case "REMOVE":
			user, err := convertEventToUser(record.Change.OldImage)
			if err != nil {
				return err
			}
			err = sendEmail(user.Email, fmt.Sprintf("Goodbye %s", user.Name), "Miss You!!!")
			if err != nil {
				return err
			}
		default:
			fmt.Println("Not correct event type")
		}

		fmt.Println("Email sent successfully")
	}
	return nil
}

func main() {
	lambda.Start(handler)
}

func sendEmail(to, body, subject string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWSREGION")),
	})
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(os.Getenv("FROMEMAILADDRESS")),
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

func convertEventToUser(event map[string]events.DynamoDBAttributeValue) (*User, error) {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)
	var user User
	for k, v := range event {

		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return nil, marshalErr
		}

		err := json.Unmarshal(bytes, &dbAttr)
		if err != nil {
			return nil, err
		}
		dbAttrMap[k] = &dbAttr
	}
	err := dynamodbattribute.UnmarshalMap(dbAttrMap, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
