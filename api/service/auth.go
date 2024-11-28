package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/xbmlz/go-web-template/api/constant"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/model"
	"github.com/xbmlz/go-web-template/api/query"
	"github.com/xbmlz/go-web-template/internal/config"
	"github.com/xbmlz/go-web-template/internal/logger"
	"github.com/xbmlz/go-web-template/internal/token"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(c *gin.Context, req dto.RegisterRequest) (resp dto.RegisterResponse, err error)
	Login(c *gin.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error)

	RefreshToken(c *gin.Context) (resp dto.LoginResponse, err error)
}

type authService struct {
}

func NewAuthService(c *config.Config) AuthService {
	return &authService{}
}

func (s *authService) Register(c *gin.Context, req dto.RegisterRequest) (resp dto.RegisterResponse, err error) {
	q := query.User
	count, err := q.Where(q.Username.Eq(req.Username)).Count()
	if err != nil {
		return resp, err
	}
	if count > 0 {
		return resp, constant.ErrUsernameExists
	}

	// create new user
	user := &model.User{}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	copier.Copy(&user, &req)

	err = q.Create(user)

	return dto.RegisterResponse{ID: user.ID}, err
}

func (s *authService) Login(c *gin.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error) {
	q := query.User
	logger.Infof("login request: %v", req)
	// check if user exists
	user, err := q.Where(q.Username.Eq(req.Username)).First()
	if err != nil {
		return resp, constant.ErrUserNotFound
	}
	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return resp, constant.ErrPasswordIncorrect
	}

	// check user status
	if user.Status != constant.UserStatusActive {
		return resp, constant.ErrUserNotActive
	}

	// generate token
	token, expiresAt, err := token.Provider.Generate(user.ID, user.Username)
	if err != nil {
		return resp, err
	}

	resp.Token = token
	resp.ExpireAt = expiresAt
	return resp, nil
}

func (s *authService) RefreshToken(c *gin.Context) (resp dto.LoginResponse, err error) {
	// get token from header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return resp, constant.ErrInvalidToken
	}

	// validate token
	claims, err := token.Provider.Validate(authHeader)
	if err != nil {
		return resp, err
	}

	// generate new token
	token, expiresAt, err := token.Provider.Generate(claims.ID, claims.Username)
	if err != nil {
		return resp, err
	}

	resp.Token = token
	resp.ExpireAt = expiresAt
	return resp, nil
}
