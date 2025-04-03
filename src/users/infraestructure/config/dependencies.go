package config

import (
	"demo/src/users/infraestructure/routers"
	"demo/src/users/application"
	"demo/src/users/infraestructure/adapters"
	"demo/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitUsers(r *gin.Engine) {
	ps, err := adapters.NewMySQL()
	if err != nil {
		panic(err)
	}


	bcryptService := adapters.NewBcryptAdapter() 

	registerUseCase := application.NewRegisterUserUseCase(ps, bcryptService)
	registerUserController := controllers.NewRegisterUserController(registerUseCase)

	updateUseCase := application.NewUpdateUserUseCase(ps)
	updateUserController := controllers. NewUUpdateUserController(updateUseCase)

	listUserUseCase := application.NewListUserUseCase(ps)
	listUserController := controllers.NewListUserController(listUserUseCase)

	deleteUserUseCase := application.NewDeleteUserUseCase(ps)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)


	routers.UserRoutes(r, registerUserController, updateUserController, listUserController, deleteUserController)
}
