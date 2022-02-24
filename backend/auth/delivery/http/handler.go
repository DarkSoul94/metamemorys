package http

import (
	"errors"
	"net/http"

	"github.com/DarkSoul94/metamemorys_backend/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC auth.AuthUC
}

func NewAuthHandler(authUC auth.AuthUC) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
	}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var user UserRegistration

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	if user.Pass != user.PassRepeat {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": errors.New("Пароли не совпадают")})
		return
	}

	token, err := h.authUC.Registration(user.Email, user.Pass, user.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success", "token": token})
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var user UserAuth

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	token, err := h.authUC.Authorization(h.toModelUser(user))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success", "token": token})
}
