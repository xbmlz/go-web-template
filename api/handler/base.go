package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/internal/validator"
)

type BaseHandler struct{}

func (h *BaseHandler) Response(c *gin.Context, data interface{}, err error) {
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

func (h *BaseHandler) ErrorWithCode(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func (h *BaseHandler) BindAndValidate(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		h.ErrorWithCode(c, http.StatusBadRequest, err)
		return err
	}

	if err := validator.Validate.Struct(obj); err != nil {
		h.ErrorWithCode(c, http.StatusBadRequest, err)
		return err
	}

	return nil
}
