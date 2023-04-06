package apiecho

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/otaxhu/bank-app/internal/api/dto"
	"github.com/otaxhu/bank-app/internal/entity"
	"github.com/otaxhu/bank-app/internal/service"
	"github.com/otaxhu/bank-app/internal/utils/encryption"
)

type echoHandlers struct {
	usersServ       service.UsersService
	dataValidator   *validator.Validate
	encryptionUtils *encryption.EncryptionUtils
}

type responseJSON struct {
	Message string `json:"message"`
}

func newResponseJSON(msg string) *responseJSON {
	return &responseJSON{Message: msg}
}

func (eh *echoHandlers) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := &dto.RegisterUser{}
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, newResponseJSON(err.Error()))
	}
	if err := eh.dataValidator.StructCtx(ctx, params); err != nil {
		return c.JSON(http.StatusBadRequest, newResponseJSON(err.Error()))
	}
	cred := &entity.UserCredentials{
		Email:    params.Email,
		Password: params.Password,
	}
	if err := eh.usersServ.RegisterUser(ctx, cred); err != nil {
		if err == service.ErrPasswordTooLong {
			return c.JSON(http.StatusBadRequest, newResponseJSON(err.Error()))
		}
		if err == service.ErrUserAlreadyRegistered {
			return c.JSON(http.StatusConflict, newResponseJSON(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, newResponseJSON(err.Error()))
	}
	return c.NoContent(http.StatusCreated)
}

func (eh *echoHandlers) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := &dto.LoginUser{}
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, newResponseJSON("Invalid Request"))
	}
	if err := eh.dataValidator.StructCtx(ctx, params); err != nil {
		return c.JSON(http.StatusBadRequest, newResponseJSON(err.Error()))
	}
	cred := &entity.UserCredentials{
		Email:    params.Email,
		Password: params.Password,
	}
	user, err := eh.usersServ.LoginUser(ctx, cred)
	if err != nil {
		if err == service.ErrPasswordTooLong {
			return c.JSON(http.StatusBadRequest, newResponseJSON(err.Error()))
		}
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusUnauthorized, newResponseJSON(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, newResponseJSON(err.Error()))
	}
	tokenString, err := eh.encryptionUtils.NewUserJWT(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, newResponseJSON(err.Error()))
	}
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
