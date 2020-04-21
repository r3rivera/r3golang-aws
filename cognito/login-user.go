package cognito

import (
	"fmt"

	aws "github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//LoginUser Authenticate the user
func (c *UserCognitoClient) LoginUser(username, password *string) UserAuthToken {
	fmt.Println("Authenticating user :: ", &username)

	m := make(map[string]*string)
	m["USERNAME"] = username
	m["PASSWORD"] = password

	rqst := &cognito.AdminInitiateAuthInput{
		AuthFlow:       aws.String(cognito.AuthFlowTypeAdminNoSrpAuth),
		ClientId:       &c.AppClientID,
		UserPoolId:     &c.UserPoolID,
		AuthParameters: m,
	}

	processor, output := c.CognitoClient.AdminInitiateAuthRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error authenticating the user! ", err)
		panic(err)
	}

	if output != nil && output.AuthenticationResult != nil {
		return UserAuthToken{
			AccessToken:  *output.AuthenticationResult.AccessToken,
			RefreshToken: *output.AuthenticationResult.RefreshToken,
			Expiration:   *output.AuthenticationResult.ExpiresIn,
		}
	}

	return UserAuthToken{}
}
