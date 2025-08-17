package bid_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/api/web/validation"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/bid_usecase"
)

type BidController struct {
	bidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
	}
}

func (u *BidController) CreateBid(c *gin.Context) {
	var bidInputDto bid_usecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDto); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.bidUseCase.CreateBid(c.Request.Context(), bidInputDto)
	if err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
