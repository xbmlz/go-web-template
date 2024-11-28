package token

import (
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xbmlz/go-web-template/api/constant"
)

type Config struct {
	Secret     string `json:"secret"`
	Expiration int64  `json:"expiration"`
}

type TokenClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type TokenProvider struct {
	config Config
}

var Provider *TokenProvider

func InitProvider(config Config) {
	Provider = &TokenProvider{config: config}
}

func (p *TokenProvider) Generate(id uint, username string) (token string, expiresAt time.Time, err error) {
	expiresAt = time.Now().Add(time.Duration(p.config.Expiration) * time.Second)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err = claims.SignedString([]byte(p.config.Secret))
	if err != nil {
		return "", expiresAt, err
	}
	return token, expiresAt, nil
}

func (p *TokenProvider) Validate(tokenString string) (claims *TokenClaims, err error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return claims, constant.ErrInvalidToken
	}

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(p.config.Secret), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return claims, constant.ErrInvalidToken
	}

	// check if token is valid
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return claims, constant.ErrInvalidToken
	}

	// check if token is expired
	if time.Unix(claims.ExpiresAt.Unix(), 0).Before(time.Now()) {
		return claims, constant.ErrTokenExpired
	}
	return claims, nil
}
