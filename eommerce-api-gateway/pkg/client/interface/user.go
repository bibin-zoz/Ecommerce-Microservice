package interfaces

import "github.com/bibin-zoz/ecommerce-api-gateway/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignUp) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
}
