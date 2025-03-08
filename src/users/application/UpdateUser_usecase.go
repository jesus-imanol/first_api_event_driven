package application

import "apiInvitation/src/users/domain/repositories"

type UpdateUserUseCase struct {
	db repositories.IUser
}

func NewUpdateUserUseCase(db repositories.IUser) *UpdateUserUseCase {
    return &UpdateUserUseCase{db: db}
}

func (uuc *UpdateUserUseCase) Execute(id int32, fullname string, email string, passwordHash string,gender string, matchPreference string, city string, state string, interests string, statusMessage string) error {
	err := uuc.db.Update(id, fullname, email, passwordHash,gender, matchPreference, city, state,interests,statusMessage)
	if err != nil {
        return err
    }
    return nil
}