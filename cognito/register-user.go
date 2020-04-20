package cognito

import (
	"fmt"

	aws "github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//RegisterUser function to signup a user of type UserCredential
func (c *UserCognitoClient) RegisterUser(user UserCredential) {
	fmt.Println("Creating new user ::", user)

	rqst := &cognito.AdminCreateUserInput{
		MessageAction: aws.String(cognito.MessageActionTypeSuppress),
		Username:      &user.Email,
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("given_name"),
				Value: aws.String(user.FirstName),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(user.FirstName),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("false"),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(user.PhoneNumber),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	fmt.Println(rqst)
	c.CognitoClient.AdminCreateUserRequest(rqst)

}
