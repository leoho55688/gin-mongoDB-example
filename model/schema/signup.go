package schema

import (
	"context"

	"backend/model/domain"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *domain.User) error
	GetUserByEmail(c context.Context, email string) (domain.User, error)
	CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error)
}
