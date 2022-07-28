package grpc

import (
	"context"
	"todoApp/model"
)

func (s *Server) SignUp(_ context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	userId, err := s.services.Authorization.CreateUser(model.User{
		Name:     req.Name,
		Username: req.UserName,
		Password: req.Password,
	})

	return &SignUpResponse{Id: int32(userId)}, err
}

func (s *Server) SignIn(_ context.Context, req *SignInRequest) (*SignInResponse, error) {
	token, err := s.services.Authorization.GenerateToken(req.UserName, req.Password)
	if err != nil {
		return &SignInResponse{
			Token:        "",
			ErrorMessage: err.Error(),
		}, err
	}

	return &SignInResponse{
		Token:        token,
		ErrorMessage: "",
	}, err
}
