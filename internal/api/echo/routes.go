package apiecho

import "github.com/labstack/echo/v4"

func (a *ApiEcho) RegisterRoutes(e *echo.Echo) {
	e.POST("/users/register", a.handlers.RegisterUser)
	e.POST("/users/login", a.handlers.LoginUser)
}
