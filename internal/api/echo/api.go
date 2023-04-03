package apiecho

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/otaxhu/bank-app/configs"
	"github.com/otaxhu/bank-app/internal/service"
)

type ApiEcho struct {
	port     string
	handlers *echoHandlers
}

func NewApiEcho(cfg *configs.Configs, us service.UsersService) *ApiEcho {
	return &ApiEcho{
		port: cfg.Port,
		handlers: &echoHandlers{
			usersServ:     us,
			dataValidator: validator.New(),
		}}
}

func (a *ApiEcho) Start(e *echo.Echo) error {
	a.RegisterRoutes(e)

	return e.Start(a.port)
}
