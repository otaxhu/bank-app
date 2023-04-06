package encryption

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/otaxhu/bank-app/internal/entity"
)

func (e *EncryptionUtils) NewUserJWT(user *entity.DomainUser) (string, error) {
	claims := &entity.UserClaims{
		Id:    user.Id,
		Roles: user.Roles,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(e.configs.JWTSecret))
}
