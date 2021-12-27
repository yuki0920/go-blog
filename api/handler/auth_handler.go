package handler

import (
	"net/http"
	"time"
	"yuki0920/go-notes/usecase"
	"yuki0920/go-notes/util"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userUsecase usecase.UserUsecase
}

type UserParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewAuthHandler(userUsecase usecase.UserUsecase) AuthHandler {
	return AuthHandler{
		userUsecase: userUsecase,
	}
}

func (authHandler *AuthHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userParam UserParam
		if err := c.Bind(&userParam); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		user, err := authHandler.userUsecase.GetByName(userParam.Name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if user.ComparePassword(userParam.Password); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		token, err := util.GenerateJwtToken(user.Name)
		if err != nil {
			return err
		}

		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 72),
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "login success",
			"token":   token,
		})

	}
}

func (authHandler *AuthHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(time.Hour * -1),
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "logout success",
		})
	}
}

type AuthOutput struct {
	IsAuthenticated bool
}

func (authHandler *AuthHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		var out AuthOutput
		out.IsAuthenticated = true

		cookie, err := c.Cookie("jwt")
		if err != nil {
			out.IsAuthenticated = false
		} else if err := util.ParseJwt(cookie.Value); err != nil {
			out.IsAuthenticated = false
		}

		return c.JSON(http.StatusOK, out)
	}
}
