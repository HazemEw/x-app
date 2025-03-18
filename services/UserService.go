package services

import (
	"errors"
	"x-app/dtos"
	"x-app/mappers"
	"x-app/models"
	"x-app/repositories"
)

func CreateUser(userDto *dtos.UserDTO) error {

	existingUser, _ := repositories.GetUserByEmail(userDto.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	cityId := uint(1)
	user := models.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
		CityID:   &cityId,
	}
	//err := SetKey("user:"+user.Email, user.Email, time.Duration(60)*time.Second)
	//if err != nil {
	//	return err
	//}
	//go func() {
	//	err := rabbitmq.PublishMessage(os.Getenv("RABBITMQ_QUEUE"), "User created: "+user.Email)
	//	if err != nil {
	//		log.Fatalf("Failed to publish message: %s", err)
	//		return
	//	}
	//}()
	return repositories.CreateUser(&user)
}

func GetAllUsers() ([]dtos.UserDTO, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}
	var usersDTO []dtos.UserDTO

	for _, user := range users {
		usersDTO = append(usersDTO, *mappers.UserToUserDTO(&user))
	}
	return usersDTO, nil
}

func GetUser(id uint) (*models.User, error) {
	var user *models.User
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	//var userDto *dtos.UserDTO = mappers.UserToUserDTO(user)

	return user, nil

}

func RemoveUser(id uint) error {
	return repositories.DeleteUser(id)

}
