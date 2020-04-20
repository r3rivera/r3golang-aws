package main

import (
	"fmt"
	"os"
	"r3golang-aws/cognito"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	c "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func awsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		fmt.Println(" AWS Session Error ::", err)
		panic(err)
	}

	return sess
}

func main() {

	session := awsSession()
	if session != nil {

		fmt.Println("Creating a new user")
		user := cognito.UserCredential{
			Email:       "rr",
			Password:    "passwordM3!",
			FirstName:   "Gopher",
			LastName:    "Tester",
			PhoneNumber: "+18007556000"}

		client := cognito.UserCognitoClient{
			CognitoClient: c.New(session),
			UserPoolID:    os.Getenv("COGNITO_USERPOOL_ID"),
			AppClientID:   os.Getenv("COGNITO_APPCLIENT_ID"),
		}
		client.RegisterUser(user)
	}

}
