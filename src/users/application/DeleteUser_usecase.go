package application

import "apiInvitation/src/users/domain/repositories"

type DeleteUserUseCase struct {
	db repositories.IUser
}

func NewDeleteUserUseCase(db repositories.IUser) *DeleteUserUseCase {
	return &DeleteUserUseCase{db: db}
}

func (duc *DeleteUserUseCase) Execute(id int32) error {
	err := duc.db.Delete(id)
    if err != nil {
        return err
    }
    return nil
}