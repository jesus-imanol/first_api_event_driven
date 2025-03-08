package repositories

import "apiInvitation/src/users/domain/entities"

type RabbitMQRepository interface {
	Publish(message *entities.User) error
}