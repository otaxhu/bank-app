package configs

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed configs.yaml
var configsFile []byte

type Configs struct {
	Port        string         `yaml:"port"`
	MysqlConfig databaseConfig `yaml:"mysql_config"`
	JWTSecret   string         `yaml:"jwt_secret"`
}

type databaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

func New() (cfg *Configs) {
	yaml.Unmarshal(configsFile, &cfg)
	return
}
