package configs

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed configs.yaml
var configsFile []byte

type Configs struct {
	Port        int            `yaml:"port"`
	MysqlConfig databaseConfig `yaml:"mysql_config"`
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
