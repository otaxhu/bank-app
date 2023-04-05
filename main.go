package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/otaxhu/bank-app/configs"
	"github.com/otaxhu/bank-app/database"
	apiecho "github.com/otaxhu/bank-app/internal/api/echo"
	"github.com/otaxhu/bank-app/internal/repository"
	"github.com/otaxhu/bank-app/internal/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			configs.New,
			database.NewMysqlConnection,
			repository.NewMysqlUsersRepository,
			service.NewUsersService,
			apiecho.NewApiEcho,
			echo.New,
		),
		fx.Invoke(
			func(apiecho *apiecho.ApiEcho, e *echo.Echo) {
				go apiecho.Start(e)
			},
			//func(ctx context.Context, us service.UsersService) {
			//	userCredentials := &entity.UserCredentials{Email: "aaabbb", Password: "12345"}
			//	anotherUser, err := us.LoginUser(ctx, userCredentials)
			//	if err != nil {
			//		log.Println(err)
			//		return
			//	}
			//	log.Println(anotherUser)
			//},
		),
	)
	app.Run()
}
