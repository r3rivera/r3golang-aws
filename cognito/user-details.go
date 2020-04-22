package cognito

import (
	"fmt"

	aws "github.com/aws/aws-sdk-go/aws"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

//IsEmailVerified Validate is the user's email is verified or not
func (c UserCognitoClient) IsEmailVerified(username *string) bool {
	fmt.Println("Getting user details! ")

	dtl := &cognito.AdminGetUserInput{
		UserPoolId: &c.UserPoolID,
		Username:   username,
	}

	userProc, userOut := c.CognitoClient.AdminGetUserRequest(dtl)
	dtlErr := userProc.Send()

	if dtlErr != nil {
		fmt.Println("Error getting user details!, ", dtlErr)
		panic(dtlErr)
	}

	isVerified := false
	for _, i := range userOut.UserAttributes {
		if i.Name == aws.String("email_verified") {
			if i.Value == aws.String("true") {
				isVerified = true
				break
			}
		}
	}

	return isVerified
}
