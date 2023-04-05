package encryption

import "github.com/otaxhu/bank-app/configs"

type EncryptionUtils struct {
	configs *configs.Configs
}

func NewEncryptionUtils(cfg *configs.Configs) *EncryptionUtils {
	return &EncryptionUtils{configs: cfg}
}
