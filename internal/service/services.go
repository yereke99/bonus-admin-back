package service

import (
	"context"

	"bonus-admin-back/config"
	"bonus-admin-back/internal/repository"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type IJWTServices interface {
	GenerateToken(email string, role string) (string, error)
	RefreshToken(tokenString string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserId(tokenString string) (string, error)
	GetCompanyId(tokenString string) (string, error)
	GetCompanyObjectId(tokenString string) (string, error)
}

type Services struct {
	JWTService IJWTServices
}

func NewServices(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repo *repository.Repositories) *Services {
	jwtServices := NewJWTService(appConfig.SecretKey, appConfig.Issuer)
	return &Services{
		JWTService: jwtServices,
	}
}
