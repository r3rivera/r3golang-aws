package cognito

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//ChangePassword Changing user password using old password with a new password
func (c *UserCognitoClient) ChangePassword(accessToken, oldPassword, newPassword *string) bool {
	log.Println("Changing user password!")

	rqst := &cognito.ChangePasswordInput{
		AccessToken:      accessToken,
		PreviousPassword: oldPassword,
		ProposedPassword: newPassword,
	}
	processor, output := c.CognitoClient.ChangePasswordRequest(rqst)
	err := processor.Send()

	if err != nil {
		log.Panicln("Error changing password!")
		panic(err)
	}
	return output != nil
}

//ForgotPassword Used to send a verification code to the user's verified email address.
func (c *UserCognitoClient) ForgotPassword(username *string) bool {
	log.Println("Forgot password flow")

	if c.IsEmailVerified(username) {
		log.Println("Email is verified! Sending forgot password!")
		rqst := &cognito.ForgotPasswordInput{
			ClientId: &c.AppClientID,
			Username: username,
		}

		processor, output := c.CognitoClient.ForgotPasswordRequest(rqst)
		err := processor.Send()

		if err != nil {
			log.Panicln("Error with the forgot password flow")
			panic(err)
		}
		return output != nil

	}

	log.Println("Email is not verified yet!")
	//Workaround is to force the email_verified attribute as true since we cannot send the verification code
	//as part of the password reset

	att := cognito.AttributeType{
		Name:  aws.String("email_verified"),
		Value: aws.String("true"),
	}

	var attrArr = make([]*cognito.AttributeType, 1)
	attrArr[0] = &att

	rqst := &cognito.AdminUpdateUserAttributesInput{
		UserPoolId:     &c.UserPoolID,
		UserAttributes: attrArr,
		Username:       username,
	}

	attrProc, _ := c.CognitoClient.AdminUpdateUserAttributesRequest(rqst)
	err := attrProc.Send()

	if err != nil {
		log.Panicln("Error updating the email verified attribute")
		panic(err)
	}

	resetRqst := &cognito.AdminResetUserPasswordInput{
		UserPoolId: &c.UserPoolID,
		Username:   username,
	}

	proc, _ := c.CognitoClient.AdminResetUserPasswordRequest(resetRqst)
	err = proc.Send()
	if err != nil {
		log.Panicln("Error reset the password!")
		panic(err)
	}

	return true
}

//ConfirmForgotPassword confirms the forgotten password initiated by the user
func (c *UserCognitoClient) ConfirmForgotPassword(username, password, code *string) bool {
	log.Println("Confirming forgotten password. Code is ", &code)

	rqst := &cognito.ConfirmForgotPasswordInput{
		Username:         username,
		Password:         password,
		ConfirmationCode: code,
		ClientId:         &c.AppClientID,
	}

	processor, output := c.CognitoClient.ConfirmForgotPasswordRequest(rqst)
	err := processor.Send()

	if err != nil {
		log.Panicln("Error confirming the code")
		panic(err)
	}
	return output != nil
}
