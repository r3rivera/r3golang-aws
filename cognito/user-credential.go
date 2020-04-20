package cognito

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//UserCredential contains information of user during signup
type UserCredential struct {
	Email       string
	Password    string
	FirstName   string
	LastName    string
	PhoneNumber string
}

//UserCognitoClient implementation
type UserCognitoClient struct {
	CognitoClient *cognito.CognitoIdentityProvider
	UserPoolID    string
	AppClientID   string
}
