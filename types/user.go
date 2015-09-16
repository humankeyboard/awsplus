package types

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

type User struct {
	UserName  string
	Password  string
	GroupName string
}

func (u *User) CreateUser(svc *iam.IAM) error {
	params := &iam.CreateUserInput{
		UserName: aws.String(u.UserName), // Required
		Path:     aws.String("/"),
	}
	_, err := svc.CreateUser(params)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) CreateLoginProfile(svc *iam.IAM) error {
	params := &iam.CreateLoginProfileInput{
		UserName:              aws.String(u.UserName),
		Password:              aws.String(u.Password),
		PasswordResetRequired: aws.Bool(true),
	}

	_, err := svc.CreateLoginProfile(params)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) AddUserToGroup(svc *iam.IAM) error {
	params := &iam.AddUserToGroupInput{
		GroupName: aws.String(u.GroupName), // Required
		UserName:  aws.String(u.UserName),  // Required
	}
	_, err := svc.AddUserToGroup(params)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser(svc *iam.IAM) error {
	params := &iam.DeleteUserInput{
		UserName: aws.String(u.UserName), // Required
	}
	_, err := svc.DeleteUser(params)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteLoginProfile(svc *iam.IAM) error {
	params := &iam.DeleteLoginProfileInput{
		UserName: aws.String(u.UserName), // Required
	}
	_, err := svc.DeleteLoginProfile(params)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) RemoveUserFromGroup(svc *iam.IAM) error {
	params := &iam.RemoveUserFromGroupInput{
		GroupName: aws.String(u.GroupName), // Required
		UserName:  aws.String(u.UserName),  // Required
	}
	_, err := svc.RemoveUserFromGroup(params)

	if err != nil {
		return err
	}

	return nil
}
