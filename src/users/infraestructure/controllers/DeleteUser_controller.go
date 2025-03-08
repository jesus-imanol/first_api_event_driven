package controllers

import (
	"apiInvitation/src/users/application"
	"net/http"
	"strconv"	
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUserUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUseCase *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{deleteUserUseCase: deleteUseCase}
}

func (du *DeleteUserController) DeleteUser (g *gin.Context) {
	idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product ID"})
        return
    }
	idDelete := int32(id)
	
	if err2 := du.deleteUserUseCase.Execute(idDelete); err2!= nil {
        g.JSON(http.StatusNotFound, gin.H{
				"detail": err2.Error(),
				"type": "users",    
		})
        return
    }
	response := gin.H{
        "data": gin.H{
            "type": "comments", 
			"id": idParam,
			"message":"User eliminado con éxito",
            
        },
    }
	g.JSON(http.StatusOK, response)
}