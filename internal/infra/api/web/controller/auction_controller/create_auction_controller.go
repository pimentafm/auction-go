package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pimentafm/auction-go/internal/infra/api/web/validation"
	"github.com/pimentafm/auction-go/internal/usecase/auction_usecase"
)

type AuctionController struct {
	AuctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		AuctionUseCase: auctionUseCase,
	}
}

func (a *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO
	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := a.AuctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
