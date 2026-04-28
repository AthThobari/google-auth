package service

import (
	"auth/model"
	"auth/repository"
	"context"
	"log"

	"google.golang.org/api/idtoken"
)

type AuthService struct {
	ClientID string
	UserRepo *repository.UserRepository
}

type GoogleUser struct {
	Email string
	Name  string
	Sub   string
	Img   string
}

func (s *AuthService) VerifyGoogleToken(ctx context.Context, token string) (*model.User, error) {

	payload, err := idtoken.Validate(ctx, token, s.ClientID)
	if err != nil {
		return nil, err
	}
	log.Print(payload)
	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	sub := payload.Claims["sub"].(string)
	picture := payload.Claims["picture"].(string)

	user, err := s.UserRepo.FindOrCreate(sub, email, name, picture)

	if err != nil {
		return nil, err
	}

	return user, nil
}
