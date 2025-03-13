package configMatch

import (
	"apiInvitation/src/match/application"
	"apiInvitation/src/match/infaestructure/adapters"
	"apiInvitation/src/match/infaestructure/controllers"
	"apiInvitation/src/match/infaestructure/routers"

	"github.com/gin-gonic/gin"
)

func InitMatch(r *gin.Engine) {
	ps, err := adapters.NewMySQL()
	if err != nil {
	panic(err)
	}
	sendMatch := application.NewSendMatchUseCase(ps)
	sendMatchController := controllers.NewSendMatchController(sendMatch)
	
	getMatchesDetails := application.NeGetUserMatchesWithDetailsUseCase(ps)
	getMatchesDetailsController := controllers.NewGetMatchesDetailsController(getMatchesDetails)

	routers.MachRouter(r,sendMatchController, getMatchesDetailsController)
}