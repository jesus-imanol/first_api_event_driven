package application

import (
	"apiInvitation/src/users/domain/entities"
	"apiInvitation/src/users/domain/repositories"
	"errors"
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
    db repositories.IUser
}

func NewLoginUserUseCase(db repositories.IUser) *LoginUserUseCase {
    return &LoginUserUseCase{db: db}
}

func (lu *LoginUserUseCase) Execute(email string, passwordHash string) (*entities.User, error) {
    users, err := lu.db.GetAll()
    if err != nil {
        return nil, err
    }

    secretKey := os.Getenv("SECRET_KEY")
    passwordWithKey := passwordHash + secretKey

    for _, user := range users {
        if user.Email == email {
            err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(passwordWithKey))
            if err != nil {
                return nil, errors.New("credenciales invalidas")
            }
            fmt.Println(user.City)
            return user, nil
        }
    }
    return nil, errors.New("user no encontrado")
}

