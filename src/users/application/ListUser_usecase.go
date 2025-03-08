package application

import (
	"apiInvitation/src/users/domain/entities"
	"apiInvitation/src/users/domain/repositories"
)

type ListUserUseCase struct {
	db repositories.IUser
}

func NewListUserUseCase(db repositories.IUser) *ListUserUseCase {
    return &ListUserUseCase{db: db}
}

func (luc *ListUserUseCase) Execute() ([]*entities.User, error) {
    users, err := luc.db.GetAll()
	if err != nil {
        return nil, err
    }
    return users, err
}