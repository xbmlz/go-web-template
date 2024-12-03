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
	"github.com/xbmlz/go-web-template/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type SysAuthService interface {
	Register(c *gin.Context, req dto.RegisterRequest) (resp dto.RegisterResponse, err error)
	Login(c *gin.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error)
	RefreshToken(c *gin.Context) (resp dto.LoginResponse, err error)
	GetUserInfo(userID uint) (resp dto.UserInfoResponse, err error)
	GetUserPermissions(userID uint) (resp dto.UserPermissionsResponse, err error)
}

type sysAuthService struct {
}

func NewSysAuthService(c *config.Config) SysAuthService {
	return &sysAuthService{}
}

func (s *sysAuthService) Register(c *gin.Context, req dto.RegisterRequest) (resp dto.RegisterResponse, err error) {
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
	req.Password = utils.HashPassword(req.Password)

	copier.Copy(&user, &req)

	err = q.Create(user)

	return dto.RegisterResponse{ID: user.ID}, err
}

func (s *sysAuthService) Login(c *gin.Context, req dto.LoginRequest) (resp dto.LoginResponse, err error) {
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
	resp.TokenPrefix = constant.TokenPrefix
	return resp, nil
}

func (s *sysAuthService) RefreshToken(c *gin.Context) (resp dto.LoginResponse, err error) {
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
	resp.TokenPrefix = constant.TokenPrefix
	return resp, nil
}

func (s *sysAuthService) GetUserInfo(userID uint) (resp dto.UserInfoResponse, err error) {
	// get user by id
	q := query.User
	user, err := q.Preload(q.Roles).FindByID(userID)
	if err != nil {
		return resp, err
	}
	// convert user to dto
	err = copier.Copy(&resp, &user)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *sysAuthService) GetUserPermissions(userID uint) (resp dto.UserPermissionsResponse, err error) {
	// get user by id
	user, err := query.User.Preload(query.User.Roles).FindByID(userID)
	if err != nil {
		return resp, err
	}
	// get role ids
	roleIDs := make([]uint, 0)
	for _, role := range user.Roles {
		roleIDs = append(roleIDs, role.ID)
	}
	// get menus by role ids
	roleMenus, err := query.RoleMenu.Where(query.RoleMenu.RoleID.In(roleIDs...)).Find()
	if err != nil {
		return resp, err
	}
	menuIDs := make([]uint, 0)
	for _, roleMenu := range roleMenus {
		menuIDs = append(menuIDs, roleMenu.MenuID)
	}
	// get menus by ids
	menus, err := query.Menu.Where(query.Menu.ID.In(menuIDs...)).Find()
	if err != nil {
		return resp, err
	}
	tree := model.BuildMenuTree(menus, 0)
	return dto.UserPermissionsResponse{
		Menus: tree,
	}, nil
}
