package controllers

import (
	"apiInvitation/src/users/application"
	"apiInvitation/src/users/domain/entities"
	"apiInvitation/src/users/infraestructure/utils"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
    registerUserUseCase *application.RegisterUserUseCase
}

func NewRegisterUserController(registerUseCase *application.RegisterUserUseCase) *RegisterUserController {
    return &RegisterUserController{registerUserUseCase: registerUseCase}
}

func GenerateJWT(user entities.User) (string, error) {
    var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userName": user.FullName,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (ruc *RegisterUserController) RegisterUser(g *gin.Context) {
    var user entities.User
    if err := g.ShouldBindJSON(&user); err != nil {
        g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }

    passwordHashed, err := utils.HashPassword(user.PasswordHash)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    registeredUser, err2 := ruc.registerUserUseCase.Execute(
        user.FullName, user.Email, passwordHashed, user.Gender,
        user.MatchPreference, user.City, user.State, user.Interests,
        user.StatusMessage, user.ProfilePicture,
    )
    if err2 != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
        return
    }

    token, err3 := GenerateJWT(user)
    if err3 != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
        return
    }

    g.Header("Authorization", "Bearer " + token)
    response := gin.H{
        "data": gin.H{
            "type": "users",
            "id":   registeredUser.Id,
            "attributes": gin.H{
                "full_name":       registeredUser.FullName,
                "email":           registeredUser.Email,
                "profile_picture": registeredUser.ProfilePicture,
                "gender":          registeredUser.Gender,
                "city":            registeredUser.City,
                "state":           registeredUser.State,
                "status_message":  registeredUser.StatusMessage,
                "match_preference": registeredUser.MatchPreference,
                "interests":       registeredUser.Interests,
            },
        },
    }
    g.JSON(http.StatusCreated, response)
}
