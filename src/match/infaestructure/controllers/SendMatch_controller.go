package controllers

import (
	"apiInvitation/src/match/application"
	"apiInvitation/src/match/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMatchContoller struct {
	sendMatchUseCase *application.SendMatchUseCase
}

func NewSendMatchController (sendMatchUseCase *application.SendMatchUseCase) *SendMatchContoller {
	return &SendMatchContoller{sendMatchUseCase: sendMatchUseCase}
}
func (sm *SendMatchContoller) SendMatch(c *gin.Context) {
	var match entities.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	macth, err := sm.sendMatchUseCase.Execute(match.SenderUser,match.ReceiverUser)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{
        "data": gin.H{
            "type": "match",
            "id":   macth.Id,
            "attributes": gin.H{
                "sender_user":       macth.SenderUser,
                "receiver_user":           macth.ReceiverUser,
            },
        },
    }
    c.JSON(http.StatusCreated, response)
}