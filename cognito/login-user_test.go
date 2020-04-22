package cognito

import (
	"os"
	"testing"

	c "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func TestLoginUser(t *testing.T) {

	os.Setenv("COGNITO_USERPOOL_ID", "")
	os.Setenv("COGNITO_APPCLIENT_ID", "")
	email := ""
	password := "passwordM3!"

	session := awsSession()
	if session != nil {

		client := UserCognitoClient{
			CognitoClient: c.New(session),
			UserPoolID:    os.Getenv("COGNITO_USERPOOL_ID"),
			AppClientID:   os.Getenv("COGNITO_APPCLIENT_ID"),
		}

		response := client.LoginUser(&email, &password)
		if response.AccessToken == "" {
			t.Error("AccessToken is expected, but value is ", response.AccessToken)
		}

		if response.RefreshToken == "" {
			t.Error("Expected RefreshToken, But value is ", response.RefreshToken)
		}

	}

}
