package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/constant"
	"github.com/xbmlz/go-web-template/internal/token"
	"github.com/xbmlz/go-web-template/internal/validator"
)

type BaseHandler struct{}

func (h *BaseHandler) Response(c *gin.Context, err error, data interface{}) {
	if err != nil {
		h.Error(c, err)
		return
	}
	h.Ok(c, data)
}

func (h *BaseHandler) Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (h *BaseHandler) Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func (h *BaseHandler) ErrorWithCode(c *gin.Context, err error, code int) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func (h *BaseHandler) BindAndValidateJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		h.ErrorWithCode(c, err, http.StatusBadRequest)
		return false
	}

	if err := validator.Validate.Struct(obj); err != nil {
		h.ErrorWithCode(c, err, http.StatusBadRequest)
		return false
	}

	return true
}

func (h *BaseHandler) BindAndValidateQuery(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBind(obj); err != nil {
		h.ErrorWithCode(c, err, http.StatusBadRequest)
		return false
	}

	if err := validator.Validate.Struct(obj); err != nil {
		h.ErrorWithCode(c, err, http.StatusBadRequest)
		return false
	}

	return true
}

func (h *BaseHandler) GetCurrentUserClaims(c *gin.Context) (claims *token.TokenClaims) {
	claimsValue, exists := c.Get(constant.CtxUserClaimsKey)
	if !exists {
		return nil
	}
	claims, exists = claimsValue.(*token.TokenClaims)
	if !exists {
		return nil
	}
	return claims
}

func (h *BaseHandler) GetCurrentUserID(c *gin.Context) (id uint) {
	claims := h.GetCurrentUserClaims(c)
	if claims == nil {
		return 0
	}
	return claims.ID
}

func (h *BaseHandler) GetCurrentUsername(c *gin.Context) (username string) {
	claims := h.GetCurrentUserClaims(c)
	if claims == nil {
		return ""
	}
	return claims.Username
}
