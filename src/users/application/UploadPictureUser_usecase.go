package application

import "apiInvitation/src/users/domain/repositories"

type UploadPictureUserUseCase struct {
	db repositories.IUser
}

func NewUploadPictureUserUseCase(db repositories.IUser) *UploadPictureUserUseCase {
    return &UploadPictureUserUseCase{db: db}
}

func (uuc *UploadPictureUserUseCase) Execute(id int32,urlPicture string) error {
	err := uuc.db.UploadPicture(id, urlPicture)
	if err != nil {
        return err
    }
    return nil
}