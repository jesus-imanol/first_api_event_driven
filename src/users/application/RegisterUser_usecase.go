package application

import (
    "apiInvitation/src/users/domain/entities"
    "apiInvitation/src/users/domain/repositories"
)

type RegisterUserUseCase struct {
    db       repositories.IUser
    rabbitMQ repositories.RabbitMQRepository
}

func NewRegisterUserUseCase(db repositories.IUser, rabbitMQ repositories.RabbitMQRepository) *RegisterUserUseCase {
    return &RegisterUserUseCase{db: db, rabbitMQ: rabbitMQ}
}

func (ru *RegisterUserUseCase) Execute(fullName, email, passwordHash, gender, matchPreference, city, state, interests, statusMessage, profilePicture string) (*entities.User, error) {
    user := entities.NewUser(fullName, email, passwordHash, gender, matchPreference, city, state, interests, statusMessage, profilePicture)
    err := ru.db.Register(user)
    if err != nil {
        return nil, err
    }
    err = ru.rabbitMQ.Publish(user)
    if err != nil {
        return nil, err
    }

    return user, nil
}
