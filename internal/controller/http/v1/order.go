package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	auth "github.com/evmartinelli/go-rifa-microservice/internal/controller/http/middleware"
	"github.com/evmartinelli/go-rifa-microservice/internal/core/order"
	"github.com/evmartinelli/go-rifa-microservice/pkg/logger"
)

type orderRoutes struct {
	useCases *UseCases
	logger   *logger.Logger
}

func newOrderRoutes(handler *gin.RouterGroup, l *logger.Logger, u *UseCases) {
	r := &orderRoutes{
		useCases: u,
		logger:   l,
	}

	h := handler.Group("/order")
	{
		h.Use(auth.JwtAuthMiddleware())
		h.POST("/", r.doPost)
		h.GET("/", r.getOrdersByUserID)
	}
}

// @Summary     Create
// @Description Create a Order
// @ID          do-post
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Param       request body order.Request true "Set up order"
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /order/ [post].
func (r *orderRoutes) doPost(c *gin.Context) {
	var request order.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - createOrder")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	res, err := r.useCases.PlaceOrder.Run(c.Request.Context(), &request, c.GetHeader("Authorization"))
	if err != nil {
		r.logger.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary     List Orders by User ID
// @Description Lists all orders from a user
// @ID          getOrdersByUserID
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Success     201 204 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /order/ [get].
func (r *orderRoutes) getOrdersByUserID(c *gin.Context) {
	res, err := r.useCases.ListOrder.Run(c.Request.Context(), c.GetHeader("Authorization"))
	if err != nil {
		r.logger.Error(err, "http - v1 - doCreateRaffle")
		errorResponse(c, http.StatusInternalServerError, "raffle service problems")

		return
	}

	if len(res) == 0 {
		c.JSON(http.StatusNoContent, res)

		return
	}

	c.JSON(http.StatusOK, res)
}
