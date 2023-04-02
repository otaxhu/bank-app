package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/bank-app/configs"
)

func NewMysqlConnection(ctx context.Context, cfg *configs.Configs) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?ParseTime=true",
		cfg.MysqlConfig.Username,
		cfg.MysqlConfig.Password,
		cfg.MysqlConfig.Host,
		cfg.MysqlConfig.Port,
		cfg.MysqlConfig.Name,
	)
	return sqlx.ConnectContext(ctx, "mysql", dsn)
}
