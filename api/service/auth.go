package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/xbmlz/go-web-template/api/constant"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/model"
	"github.com/xbmlz/go-web-template/api/query"
	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

type AuthService interface {
	Register(c *gin.Context, req dto.RegisterRequest) (resp dto.RegisterResponse, err error)
}

func NewAuthService() AuthService {
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
