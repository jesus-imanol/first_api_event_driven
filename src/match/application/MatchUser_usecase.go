package application

import (
	"apiInvitation/src/match/domain/entities"
	"apiInvitation/src/match/domain/repositories"
)

type SendMatchUseCase struct {
	matchRepository repositories.IMatch
}
func NewSendMatchUseCase(matchRepository repositories.IMatch) *SendMatchUseCase {
	return &SendMatchUseCase{matchRepository: matchRepository}
}
func (sm *SendMatchUseCase) Execute(senderUser int32,receiverUser int32) (*entities.Match,error) {
	match := entities.NewMatch(senderUser,receiverUser)
	err := sm.matchRepository.Send(match)

	if err!= nil {
		return nil,err
	}
	return match,err
}