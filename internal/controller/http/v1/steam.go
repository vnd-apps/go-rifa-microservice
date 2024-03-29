package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type steamRoutes struct {
	useCases *UseCases
	l        *logger.Logger
}

func newSteamRoutes(handler *gin.RouterGroup, l *logger.Logger, u *UseCases) {
	r := &steamRoutes{u, l}

	h := handler.Group("/steam")
	{
		h.POST("/do-player-inventory", r.doPlayerInventory)
	}
}

type doSteamRequest struct {
	SteamID string `json:"steam_id" binding:"required" example:"894012849024820948209"`
}

// @Summary     Create
// @Description Create a Player Inventory
// @ID          do-player-inventory
// @Tags  	    steam
// @Accept      json
// @Produce     json
// @Param       request body doSteamRequest true "set up steam"
// @Success     200 {object} skin.Skin
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /steam/do-player-inventory [post].
func (r *steamRoutes) doPlayerInventory(c *gin.Context) {
	var request doSteamRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	skins, err := r.useCases.PlayerInventory.Run(
		c.Request.Context(),
		request.SteamID,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	c.JSON(http.StatusOK, skins)
}
