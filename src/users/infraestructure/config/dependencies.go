package config

import (
	"apiInvitation/src/users/infraestructure/routers"
	"apiInvitation/src/users/application"
	"apiInvitation/src/users/infraestructure/adapters"
	"apiInvitation/src/users/infraestructure/controllers"
	"os"
	"github.com/gin-gonic/gin"
	"fmt"
)


func InitUsers(r *gin.Engine) {
	ps, err := adapters.NewMySQL()
	if err != nil {
	panic(err)
	}
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
    rabbitmqPass := os.Getenv("RABBITMQ_PASS")
    rabbitmqHost := os.Getenv("RABBITMQ_HOST")
    rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)

    rabbitMQAdapter := adapters.NewRabbitMQAdapter(connStr)
	registerUseCase := application.NewRegisterUserUseCase(ps, rabbitMQAdapter)
	registerUser_controller := controllers.NewRegisterUserController(registerUseCase)

	updateUseCase := application.NewUpdateUserUseCase(ps)
	updateUser_controller := controllers.NewUpdateUserController(updateUseCase)
	// upload file picturee
	uploadPictureUseCase := application.NewUploadPictureUserUseCase(ps)
	uploadPictureUser_controller := controllers.NewUploadPictureUserController(uploadPictureUseCase)


	listUserUseCase := application.NewListUserUseCase(ps)
	listUser_controller := controllers.NewListUserController(listUserUseCase)

	// DELETE USER
	deleteUserUseCase := application.NewDeleteUserUseCase(ps)
	deleteUser_controller := controllers.NewDeleteUserController(deleteUserUseCase)

	loginUseCase := application.NewLoginUserUseCase(ps)
    loginUser_controller := controllers.NewLoginUserController(loginUseCase)

	getByIdUseCase := application.NewGetUserById(ps)
	getUserById_controller := controllers.NewGetUserByIDController(getByIdUseCase)

	routers.UserRoutes(r,registerUser_controller, updateUser_controller, listUser_controller, deleteUser_controller, loginUser_controller, getUserById_controller, uploadPictureUser_controller)
}