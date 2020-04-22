package cognito

import (
	"os"
	"testing"

	c "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func TestIsEmailNotVerified(t *testing.T) {

	session := awsSession()
	if session != nil {

		client := UserCognitoClient{
			CognitoClient: c.New(session),
			UserPoolID:    os.Getenv("COGNITO_USERPOOL_ID"),
			AppClientID:   os.Getenv("COGNITO_APPCLIENT_ID"),
		}

		email := ""
		response := client.IsEmailVerified(&email)

		if response {
			t.Error("Expected to be TRUE, But value is ", response)
		}
	}

}
