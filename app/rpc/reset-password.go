package rpc

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/clintvidler/identity-go/app/data"
	proto "github.com/clintvidler/identity-go/gen/proto/server"
)

func (s IdentityService) ResetPasswordStart(ctx context.Context, req *proto.ResetPasswordStartRequest) (*proto.ResetPasswordStartReponse, error) {
	// Required fields
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	// Email should exist
	_, err := s.data.User.ReadOne(0, req.Email)
	if err != nil {
		return nil, err
	}

	// Store pending confirmation record for the email
	key, err := s.data.User.UpsertPendingResetPassword(req.Email)
	if err != nil {
		return nil, err
	}

	emailTo := req.Email
	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Reset your password"
	emailBody := "http://" + os.Getenv("FRONTEND_URL") + "/login/reset/" + key + "\n\nExpires in 24 hours."
	s.emailClient.Send(emailTo, emailFrom, emailSubject, emailBody)

	return &proto.ResetPasswordStartReponse{}, nil
}

func (s IdentityService) ResetPasswordPending(ctx context.Context, req *proto.ResetPasswordPendingRequest) (*proto.ResetPasswordPendingReponse, error) {
	email, err := s.data.User.ReadPendingResetPassword(req.Key)
	if err != nil {
		return nil, err
	}

	return &proto.ResetPasswordPendingReponse{Email: email}, nil
}

func (s IdentityService) ResetPasswordFinish(ctx context.Context, req *proto.ResetPasswordFinishRequest) (*proto.ResetPasswordFinishReponse, error) {
	// Required fields
	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	email, err := s.data.User.ReadPendingResetPassword(req.Key)
	if err != nil {
		return nil, err
	}

	var user data.User
	user.Email = email
	user.SetPassword(req.Password)

	err = s.data.User.ResetPassword(user)
	if err != nil {
		return nil, err
	}

	emailTo := email
	emailFrom := "no-reply@" + strings.Split(os.Getenv("FRONTEND_URL"), ":")[0]
	emailSubject := "Welcome, account confirmed"
	emailBody := "http://" + os.Getenv("FRONTEND_URL") + "/login"

	s.emailClient.Send(emailTo, emailFrom, emailSubject, emailBody)

	return &proto.ResetPasswordFinishReponse{}, nil
}
