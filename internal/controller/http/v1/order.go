package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type orderRoutes struct {
	useCases UseCases
	logger   logger.Interface
}

func newOrderRoutes(handler *gin.RouterGroup, l logger.Interface, u UseCases) {
	r := &orderRoutes{
		useCases: u,
		logger:   l,
	}

	h := handler.Group("/raffle")
	{
		h.GET("/", r.getOrder)
		h.POST("/", r.doPost)
	}
}

// @Summary     Show raffles
// @Description Show all available raffles
// @ID          available
// @Tags  	    raffle
// @Accept      json
// @Produce     json
// @Success     200 {object} availableResponse
// @Failure     500 {object} response
// @Router      /order/ [get].
func (r *orderRoutes) getOrder(c *gin.Context) {
	c.JSON(http.StatusOK, availableResponse{})
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
// @Router      /order/ [post].
func (r *orderRoutes) doPost(c *gin.Context) {
	var request order.PlaceOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - createOrder")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	res, err := r.useCases.PlaceOrder.Run(c.Request.Context(), &request)
	if err != nil {
		r.logger.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	c.JSON(http.StatusCreated, res)
}
