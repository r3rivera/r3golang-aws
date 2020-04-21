package cognito

import (
	"fmt"

	aws "github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//GenerateVerficationCodeForEmail generates code for email with the AWS AccessToken receieved during authentication
func (c *UserCognitoClient) GenerateVerficationCodeForEmail(accessToken *string) bool {
	fmt.Println("Generating code for verifying email address using accessToken :: ", &accessToken)

	rqst := &cognito.GetUserAttributeVerificationCodeInput{
		AccessToken:   accessToken,
		AttributeName: aws.String("email"),
	}

	processor, output := c.CognitoClient.GetUserAttributeVerificationCodeRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error generating code for email verification!", err)
		panic(err)
	}
	return output != nil
}

//VerifyEmailByCode verfies the email address using the accessToken
func (c *UserCognitoClient) VerifyEmailByCode(accessToken, code *string) bool {
	fmt.Println("Verifying email address using accessToken and code!")

	rqst := &cognito.VerifyUserAttributeInput{
		AccessToken:   accessToken,
		AttributeName: aws.String("email"),
		Code:          code,
	}

	processor, output := c.CognitoClient.VerifyUserAttributeRequest(rqst)
	err := processor.Send()

	if err != nil {
		fmt.Println("Error verifying email using code!", err)
		panic(err)
	}
	return output != nil
}
