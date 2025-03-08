package controllers

import (
	"apiInvitation/src/users/application"
	"apiInvitation/src/users/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUserUseCase *application.UpdateUserUseCase
}

func NewUpdateUserController(updateUseCase * application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{updateUserUseCase: updateUseCase}
}

func (uuc *UpdateUserController) UpdateUser(g *gin.Context) {
	var user entities.User
    idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
    if err!= nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := g.ShouldBindJSON(&user); err!= nil {
        g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	idCheck := int32(id)

    err = uuc.updateUserUseCase.Execute(idCheck, user.FullName, user.Email, user.PasswordHash, user.Gender,
        user.MatchPreference, user.City, user.State, user.Interests,
        user.StatusMessage,)
    if err!= nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    response := gin.H{
        "data": gin.H{
            "type": "users",
            "id":   user.Id,
            "attributes": gin.H{
                "full_name":       user.FullName,
                "email":           user.Email,
                "gender":          user.Gender,
                "city":            user.City,
                "state":           user.State,
                "status_message":  user.StatusMessage,
                "match_preference": user.MatchPreference,
                "interests":       user.Interests,
            },
        },
    }
    g.JSON(http.StatusOK, response)
}
