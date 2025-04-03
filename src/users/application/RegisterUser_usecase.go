package application

import (
	"fmt"
	"demo/src/users/domain/entities"
	"demo/src/users/domain/repositories"
	"demo/src/users/domain/services"
)

type RegisterUserUseCase struct {
	UserRepository repositories.IUser
	BcryptService  services.IBcryptService
}

func NewRegisterUserUseCase(userRepository repositories.IUser, bcryptService services.IBcryptService) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepository,
		BcryptService:  bcryptService,
	}
}

func (ru *RegisterUserUseCase) Execute(name string, lastname string, password string, role int32) (*entities.User, error) {
	if name == "" || lastname == "" || password == "" {
		return nil, fmt.Errorf("todos los campos son obligatorios")
	}


	hashedPassword, err := ru.BcryptService.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error al encriptar la contraseña: %v", err)
	}

	user := entities.NewUser(name, lastname, hashedPassword, role)


	err = ru.UserRepository.Register(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
