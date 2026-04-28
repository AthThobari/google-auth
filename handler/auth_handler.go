package handler

import (
	"context"

	"auth/pb"
	"auth/pkg"
	"auth/repository"
	"auth/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	Service   *service.AuthService
	JWTSecret string
	UserRepo  *repository.UserRepository
}

func (h *AuthHandler) LoginWithGoogle(ctx context.Context, req *pb.GoogleLoginRequest) (*pb.AuthResponse, error) {

	user, err := h.Service.VerifyGoogleToken(ctx, req.IdToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid google token")
	}

	if user.Role == "" {
		user.Role = "user"
	}

	token, err := pkg.GenerateJWT(user.ID, user.Name, user.Email, user.Role, h.JWTSecret)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed generate jwt")
	}

	return &pb.AuthResponse{
		AccessToken: token,
		Email:       user.Email,
		Username:        user.Name,
		Role: user.Role,
		UserId: user.ID,
	}, nil
}
