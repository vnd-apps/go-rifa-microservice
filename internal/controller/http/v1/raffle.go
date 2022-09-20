package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/raffle"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type raffleRoutes struct {
	useCases *UseCases
	logger   logger.Interface
}

func newRaffleRoutes(handler *gin.RouterGroup, l logger.Interface, u *UseCases) {
	r := &raffleRoutes{
		useCases: u,
		logger:   l,
	}

	h := handler.Group("/raffle")
	{
		h.GET("/", r.getAll)
		h.GET("/:id", r.getbyID)
		h.POST("/", r.doCreateRaffle)
	}
}

type availableResponse struct {
	Available []raffle.Raffle `json:"data"`
}

// @Summary     Show raffles
// @Description Show all available raffles
// @ID          getAll
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Success     200 {object} availableResponse
// @Failure     500 {object} response
// @Router      /raffle/ [get].
func (r *raffleRoutes) getAll(c *gin.Context) {
	raffles, err := r.useCases.ListRaffle.Run(c.Request.Context())
	if err != nil {
		r.logger.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, availableResponse{raffles})
}

// @Summary     Show raffles
// @Description Show raffle by ID
// @ID          getbyID
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Success     200 {object} availableResponse
// @Failure     500 {object} response
// @Router      /raffle/:id [get].
func (r *raffleRoutes) getbyID(c *gin.Context) {
	raffleResponse, err := r.useCases.GetRaffle.Run(c.Request.Context(), c.Param("id"))
	if err != nil {
		r.logger.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, raffleResponse)
}

// @Summary     Create
// @Description Create a Raffle
// @ID          do-create
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Param       request body raffle.Request true "Set up raffle"
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /raffle/ [post].
func (r *raffleRoutes) doCreateRaffle(c *gin.Context) {
	var request raffle.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.useCases.GenerateRaffle.Run(
		c.Request.Context(),
		&raffle.Raffle{
			Name:        request.Name,
			Description: request.Description,
			ImageURL:    request.ImageURL,
			UnitPrice:   request.UnitPrice,
			Quantity:    request.Quantity,
		},
	)
	if err != nil {
		r.logger.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	c.Status(http.StatusCreated)
}
