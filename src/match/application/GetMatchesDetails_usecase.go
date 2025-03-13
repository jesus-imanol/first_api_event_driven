package application

import (
	"apiInvitation/src/match/domain/models"
	"apiInvitation/src/match/domain/repositories"
)

type GetUserMatchesWithDetailsUseCase struct {
	db repositories.IMatch
}
func NeGetUserMatchesWithDetailsUseCase (db repositories.IMatch) *GetUserMatchesWithDetailsUseCase {
	return &GetUserMatchesWithDetailsUseCase{db: db}
}

func (ru *GetUserMatchesWithDetailsUseCase) Execute(userId int32) ([]*models.MatchWithDetails, error) {
	matches, err := ru.db.GetUserMatchesWithDetails(userId)
	if err!= nil {
		return nil, err
	}	
	return matches, nil
}