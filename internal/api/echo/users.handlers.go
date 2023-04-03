package apiecho

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/otaxhu/bank-app/internal/api/dto"
	"github.com/otaxhu/bank-app/internal/entity"
	"github.com/otaxhu/bank-app/internal/service"
)

type echoHandlers struct {
	usersServ     service.UsersService
	dataValidator *validator.Validate
}

type responseMessage struct {
	Message string `json:"message"`
}

func (eh *echoHandlers) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	cred := &dto.RegisterUser{}
	if err := c.Bind(cred); err != nil {
		return c.JSON(http.StatusBadRequest, &responseMessage{Message: "Invalid Request"})
	}
	if err := eh.usersServ.RegisterUser(ctx, &entity.UserCredentials{
		Email:    cred.Email,
		Password: cred.Password,
	}); err != nil {
		if err == service.ErrUserAlreadyRegistered {
			return c.JSON(http.StatusConflict, &responseMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, &responseMessage{Message: err.Error()})
	}
	return c.NoContent(http.StatusCreated)
}
