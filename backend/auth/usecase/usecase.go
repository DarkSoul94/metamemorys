package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/DarkSoul94/metamemorys_backend/auth"
	"github.com/DarkSoul94/metamemorys_backend/models"
	"github.com/DarkSoul94/metamemorys_backend/pkg/logger"
	"github.com/dgrijalva/jwt-go/v4"
)

var (
	ErrInvalidUser = errors.New("Не правильный email или пароль")
	ErrAuthError   = errors.New("Ошибка авторизации пользователя")
)

type authUC struct {
	authRP         auth.AuthRepo
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUC(authRP auth.AuthRepo,
	signingKey []byte,
	tokenTTL time.Duration) auth.AuthUC {
	return &authUC{
		authRP:         authRP,
		signingKey:     signingKey,
		expireDuration: tokenTTL,
	}
}

func (u *authUC) Registration(email, pass, name string) (string, error) {
	var newUser *models.User = &models.User{
		Name:  name,
		Email: email,
	}

	if hash, err := HashPassword(pass); err == nil {
		newUser.PassHash = hash
	} else {
		return "", err
	}

	if err := u.authRP.CreateUser(newUser); err != nil {
		if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
			return "", errors.New("Пользователь с таким email уже существует")
		}

		logger.LogError(
			"Registration",
			"auth/usecase.go",
			fmt.Sprintf("email: %s; pass: %s; name: %s", email, pass, name),
			err,
		)
		return "", err
	}

	existUser, err := u.authRP.GetUser(email)
	if err != nil {
		return "", err
	}

	return u.generateToken(existUser)
}

func (u *authUC) Authorization(user *models.User) (string, error) {
	existUser, err := u.authRP.GetUser(user.Email)
	if err != nil || !CheckPasswordHash(user.PassHash, existUser.PassHash) {
		return "", ErrInvalidUser
	}

	return u.generateToken(existUser)
}

type claimsUser struct {
	ID    uint64
	Name  string
	Email string
}

type authClaims struct {
	jwt.StandardClaims
	User *claimsUser `json:"user"`
}

func (u *authUC) generateToken(user *models.User) (string, error) {
	var (
		token    *jwt.Token
		strToken string
		err      error
	)

	claims := authClaims{
		User: &claimsUser{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(u.expireDuration * time.Second)),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err = token.SignedString(u.signingKey)
	if err != nil {
		logger.LogError("GenerateToken", "auth/usecase", user.Email, err)
		return "", ErrAuthError
	}

	return strToken, nil
}

func (u *authUC) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return u.signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authClaims); ok && token.Valid {
		return &models.User{
			ID:       claims.User.ID,
			Name:     claims.User.Name,
			Email:    claims.User.Email,
			PassHash: "",
		}, nil
	}

	return nil, nil
}
