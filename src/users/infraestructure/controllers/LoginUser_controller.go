package controllers

import (
    "apiInvitation/src/users/application"
    "apiInvitation/src/users/domain/entities"
    "net/http"
    "github.com/gin-gonic/gin"
)

type LoginUserController struct {
    loginUserUseCase *application.LoginUserUseCase
}

func NewLoginUserController(loginUseCase *application.LoginUserUseCase) *LoginUserController {
    return &LoginUserController{loginUserUseCase: loginUseCase}
}

func (luc *LoginUserController) LoginUser(g *gin.Context) {
    var user entities.User
    if err := g.ShouldBindJSON(&user); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }

    foundUser, err := luc.loginUserUseCase.Execute(user.Email, user.PasswordHash)
    if err != nil {
        g.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    token, err := GenerateJWT(*foundUser)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    g.Header("Authorization", "Bearer " + token)

    response := gin.H{
        "data": gin.H{
            "type": "users",
            "id":   foundUser.Id,
            "attributes": gin.H{
                "full_name":     foundUser.FullName,
                "profile_picture":    foundUser.ProfilePicture,
                "gender":     foundUser.Gender,
                "city":     foundUser.State,
                "status_message" : foundUser.StatusMessage,
                "match_preference": foundUser.MatchPreference,	
                "interests": foundUser.Interests,
                "email": foundUser.Email,
                "state": foundUser.State,
            },
        },
    }
    g.JSON(http.StatusOK, response)
}
