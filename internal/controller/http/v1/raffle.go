package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
	"github.com/evmartinelli/go-rifa-microservice/internal/usecase"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type raffleRoutes struct {
	t usecase.Raffle
	l logger.Interface
}

func newRaffleRoutes(handler *gin.RouterGroup, t usecase.Raffle, l logger.Interface) {
	r := &raffleRoutes{t, l}

	h := handler.Group("/raffle")
	{
		h.GET("/avaliable", r.history)
		h.POST("/do-create", r.doCreateRaffle)
	}
}

type avaliableResponse struct {
	Avaliable []entity.Raffle `json:"avaliable"`
}

// @Summary     Show raffles
// @Description Show all avaliable raffles
// @ID          raffle
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} raffleResponse
// @Failure     500 {object} response
// @Router      /raffle/avaliable [get]
func (r *raffleRoutes) history(c *gin.Context) {
	raffles, err := r.t.GetAvaliableRaffle(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, avaliableResponse{raffles})
}

type doRaffleRequest struct {
	Name         string `json:"name" binding:"required" example:"rifa faca"`
	Value        int    `json:"value" binding:"required" example:"1"`
	TotalNumbers int    `json:"totalnumbers" binding:"required" example:"20"`
}

// @Summary     Create
// @Description Create a Raffle
// @ID          do-create
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Param       request body doRaffleRequest true "Set up raffle"
// @Success     200 {object} entity.Raffle
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /avaliable/do-create [post]
func (r *raffleRoutes) doCreateRaffle(c *gin.Context) {
	var request doRaffleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Create(
		c.Request.Context(),
		entity.Raffle{
			Name:         request.Name,
			TotalNumbers: request.TotalNumbers,
			Value:        request.Value,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	c.JSON(http.StatusCreated, "")
}
