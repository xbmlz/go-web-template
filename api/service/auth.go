package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/dto"
)

type authService struct{}

type AuthService interface {
	Register(c *gin.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Register(c *gin.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// TODO: implement register logic
	return dto.RegisterResponse{}, nil
}
