package cognito

import (
	"fmt"
	"os"
	"testing"

	aws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	c "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func awsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	if err != nil {
		fmt.Println(" AWS Session Error ::", err)
		panic(err)
	}

	return sess
}

func TestRegisterUser(t *testing.T) {
	os.Setenv("COGNITO_USERPOOL_ID", "")
	os.Setenv("COGNITO_APPCLIENT_ID", "")

	session := awsSession()
	if session != nil {

		fmt.Println("Creating a new user")
		user := UserCredential{
			Email:       "",
			Password:    "passwordM3!",
			FirstName:   "Gopher",
			LastName:    "Tester",
			PhoneNumber: "+18007556000"}

		client := UserCognitoClient{
			CognitoClient: c.New(session),
			UserPoolID:    os.Getenv("COGNITO_USERPOOL_ID"),
			AppClientID:   os.Getenv("COGNITO_APPCLIENT_ID"),
		}

		response := client.RegisterUser(user)
		if response.AccessToken == "" {
			t.Error("Expected AccessToken, But value is ", response.AccessToken)
		}

		if response.RefreshToken == "" {
			t.Error("Expected RefreshToken, But value is ", response.RefreshToken)
		}
	}
}
