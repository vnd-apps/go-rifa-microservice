package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/internal/usecase"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type steamRoutes struct {
	t usecase.Steam
	l logger.Interface
}

func newSteamRoutes(handler *gin.RouterGroup, t usecase.Steam, l logger.Interface) {
	r := &steamRoutes{t, l}

	h := handler.Group("/steam")
	{
		h.POST("/do-player-inventory", r.doplayerInventory)
	}
}

type doSteamRequest struct {
	SteamID string `json:"steam_id" binding:"required" example:"894012849024820948209"`
}

// @Summary     Create
// @Description Create a Raffle
// @ID          do-create
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Param       request body doRaffleRequest true "Set up raffle"
// @Success     200 {object} entity.Skin
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /steam//do-player-inventory [post].
func (r *steamRoutes) doplayerInventory(c *gin.Context) {
	var request doSteamRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	skin, err := r.t.GetPlayerInventory(
		c.Request.Context(),
		request.SteamID,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	c.JSON(http.StatusOK, skin)
}
