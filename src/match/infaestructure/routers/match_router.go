package routers

import (
	"apiInvitation/src/match/infaestructure/controllers"

	"github.com/gin-gonic/gin"
)

func MachRouter(r *gin.Engine, sendMatchController *controllers.SendMatchContoller, getMatchesDetailsController *controllers.GetMatchesDetailsController) {
	v1 := r.Group("/v1/match")
	{
		v1.POST("/sendMatch", sendMatchController.SendMatch)
		v1.GET("/getMatchesDetails/:id", getMatchesDetailsController.GetMatchesDetails)

	}
}