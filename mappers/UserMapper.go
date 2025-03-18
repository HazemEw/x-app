package mappers

import (
	"x-app/dtos"
	"x-app/models"
)

func UserToUserDTO(user *models.User) *dtos.UserDTO {
	return &dtos.UserDTO{
		Name:  user.Name,
		Email: user.Email,
	}
}
