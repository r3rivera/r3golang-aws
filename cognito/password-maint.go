package cognito

import (
	"fmt"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//ChangePassword Changing user password using old password with a new password
func (c *UserCognitoClient) ChangePassword(accessToken, oldPassword, newPassword *string) bool {
	fmt.Println("Changing user password!")

	rqst := &cognito.ChangePasswordInput{
		AccessToken:      accessToken,
		PreviousPassword: oldPassword,
		ProposedPassword: newPassword,
	}
	processor, output := c.CognitoClient.ChangePasswordRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error changing password!")
		panic(err)
	}
	return output != nil
}

//ForgotPassword Used to send a verification code to the user's verified email address.
func (c *UserCognitoClient) ForgotPassword(username *string) bool {
	fmt.Println("Forgot password flow")

	rqst := &cognito.ForgotPasswordInput{
		ClientId: &c.AppClientID,
		Username: username,
	}

	processor, output := c.CognitoClient.ForgotPasswordRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error with the forgot password flow")
		panic(err)
	}
	return output != nil
}

//ConfirmForgotPassword confirms the forgotten password initiated by the user
func (c *UserCognitoClient) ConfirmForgotPassword(username, password, code *string) bool {
	fmt.Println("Confirming forgotten password. Code is ", &code)

	rqst := &cognito.ConfirmForgotPasswordInput{
		Username:         username,
		Password:         password,
		ConfirmationCode: code,
		ClientId:         &c.AppClientID,
	}

	processor, output := c.CognitoClient.ConfirmForgotPasswordRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error confirming the code")
		panic(err)
	}
	return output != nil
}
