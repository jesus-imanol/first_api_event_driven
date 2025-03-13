package controllers

import (
	"apiInvitation/src/match/application"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type GetMatchesDetailsController struct {
	getMatchesDetailsUseCase *application.GetUserMatchesWithDetailsUseCase
}
func NewGetMatchesDetailsController(getMatchesDetailsUse *application.GetUserMatchesWithDetailsUseCase) *GetMatchesDetailsController {
	return &GetMatchesDetailsController{getMatchesDetailsUseCase: getMatchesDetailsUse}
}

func (c *GetMatchesDetailsController) GetMatchesDetails(g *gin.Context) {
	idParam := g.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid product ID"})
        return
    }
	idGet:= int32(id)
	matches, err := c.getMatchesDetailsUseCase.Execute(idGet)
	if err!= nil {
        g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
	g.JSON(http.StatusOK, matches)
}