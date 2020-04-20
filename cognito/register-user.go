package cognito

import (
	"fmt"

	aws "github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//RegisterUser function to signup a user of type UserCredential without verifying the password.
//This will have a status of FORCE_CHANGE_PASSWORD
func (c *UserCognitoClient) RegisterUser(user UserCredential) UserAuthToken {

	rqst := &cognito.AdminCreateUserInput{
		MessageAction:     aws.String(cognito.MessageActionTypeSuppress),
		Username:          &user.Email,
		UserPoolId:        &c.UserPoolID,
		TemporaryPassword: aws.String(user.Password),
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

	processor, output := c.CognitoClient.AdminCreateUserRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error sending the request! ", err)
		panic(err)
	}

	//UUID of the created usr
	if output.User.Username != nil {
		return c.adminInitiate(&user.Email, &user.Password)
	}
	return UserAuthToken{}
}

func (c *UserCognitoClient) adminInitiate(user, password *string) UserAuthToken {

	m := make(map[string]*string)
	m["USERNAME"] = user
	m["PASSWORD"] = password

	rqst := &cognito.AdminInitiateAuthInput{
		ClientId:       &c.AppClientID,
		UserPoolId:     &c.UserPoolID,
		AuthFlow:       aws.String(cognito.AuthFlowTypeAdminNoSrpAuth),
		AuthParameters: m,
	}
	processor, output := c.CognitoClient.AdminInitiateAuthRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error with admin initiate user! ", err)
		panic(err)
	}

	if output != nil {

		challenge := *output.ChallengeName
		if challenge == "NEW_PASSWORD_REQUIRED" && output.Session != nil {
			return c.adminRespondToChallenge(user, password, output.Session)
		}
	}
	return UserAuthToken{}
}

func (c *UserCognitoClient) adminRespondToChallenge(user, password, session *string) UserAuthToken {
	m := make(map[string]*string)
	m["USERNAME"] = user
	m["NEW_PASSWORD"] = password

	rqst := &cognito.AdminRespondToAuthChallengeInput{
		ClientId:           &c.AppClientID,
		UserPoolId:         &c.UserPoolID,
		ChallengeName:      aws.String(cognito.ChallengeNameTypeNewPasswordRequired),
		Session:            session,
		ChallengeResponses: m,
	}

	processor, output := c.CognitoClient.AdminRespondToAuthChallengeRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error with responding to admin initiate user! ", err)
		panic(err)
	}

	if output != nil {
		return UserAuthToken{
			AccessToken:  *output.AuthenticationResult.AccessToken,
			RefreshToken: *output.AuthenticationResult.RefreshToken,
			Expiration:   *output.AuthenticationResult.ExpiresIn,
		}
	}
	return UserAuthToken{}
}
