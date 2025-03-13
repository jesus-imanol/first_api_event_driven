package repositories

import (
	"apiInvitation/src/match/domain/entities"
	"apiInvitation/src/match/domain/models"
)

type IMatch interface {
	Send(match *entities.Match) error
	GetUserMatchesWithDetails(userId int32) ([]*models.MatchWithDetails, error)
}