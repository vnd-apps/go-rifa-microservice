package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type orderRoutes struct {
	useCases *UseCases
	logger   logger.Interface
}

func newOrderRoutes(handler *gin.RouterGroup, l logger.Interface, u *UseCases) {
	r := &orderRoutes{
		useCases: u,
		logger:   l,
	}

	h := handler.Group("/order")
	{
		h.POST("/", r.doPost)
	}
}

// @Summary     Create
// @Description Create a Order
// @ID          do-post
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Param       request body order.PlaceOrderRequest true "Set up order"
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
