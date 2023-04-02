package main

import (
	"context"

	"github.com/otaxhu/bank-app/configs"
	"github.com/otaxhu/bank-app/database"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			configs.New,
			database.NewMysqlConnection,
		),
		fx.Invoke(),
	)
	app.Run()
}
