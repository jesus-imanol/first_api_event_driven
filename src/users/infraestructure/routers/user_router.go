package routers

import (
    "apiInvitation/src/users/infraestructure/controllers"
    "apiInvitation/src/users/infraestructure/middleware"
    "github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, registerController *controllers.RegisterUserController, updateController *controllers.UpdateUserController, listUserController *controllers.ListUserController, deleteController *controllers.DeleteUserController, loginUserController *controllers.LoginUserController, getUserByIdController *controllers.GetUserByIdController, uploadPictureController *controllers.UploadPictureUserController) {
    v1 := r.Group("/v1/users")
    {
        v1.POST("/", registerController.RegisterUser)
        v1.POST("/login", loginUserController.LoginUser)
        v1.GET("/:id",getUserByIdController.GetUserByID)
        v1.GET("/all", listUserController.GetAllUsers)
        v1.PUT("/upload-picture/:id", uploadPictureController.UpdatePictureUser)
       

    }

    v1Auth := r.Group("/v1/users")
    v1Auth.Use(middleware.AuthMiddleware())
    {
        v1Auth.PUT("/:id", updateController.UpdateUser)
        v1Auth.GET("/", listUserController.GetAllUsers)
        v1Auth.DELETE("/:id", deleteController.DeleteUser)
    }
}
