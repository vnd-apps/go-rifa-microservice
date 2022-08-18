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
		h.GET("/available", r.available)
		h.POST("/do-create", r.doCreateRaffle)
	}
}

type availableResponse struct {
	Available []entity.Raffle `json:"available"`
}

// @Summary     Show raffles
// @Description Show all available raffles
// @ID          available
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Success     200 {object} availableResponse
// @Failure     500 {object} response
// @Router      /raffle/available [get].
func (r *raffleRoutes) available(c *gin.Context) {
	raffles, err := r.t.GetAvailableRaffle(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, availableResponse{raffles})
}

type doRaffleRequest struct {
	Name         string `json:"name" binding:"required" example:"rifa faca"`
	Value        int    `json:"value" binding:"required" example:"1"`
	TotalNumbers int    `json:"totalNumbers" binding:"required" example:"20"`
}

// @Summary     Create
// @Description Create a Raffle
// @ID          do-create
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Param       request body doRaffleRequest true "Set up raffle"
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /available/do-create [post].
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

	c.Status(http.StatusCreated)
}
