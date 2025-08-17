package auction_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/rest_err"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/api/web/validation"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/auction_usecase"
)

type auctionController struct {
	auctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *auctionController {
	return &auctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (u *auctionController) CreateAuction(c *gin.Context) {
	var auctionInputDto auction_usecase.AuctionInputDto

	if err := c.ShouldBindJSON(&auctionInputDto); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.auctionUseCase.CreateAuction(c.Request.Context(), auctionInputDto)
	if err != nil {
		restErr := rest_err.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
