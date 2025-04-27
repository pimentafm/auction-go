package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pimentafm/auction-go/internal/infra/api/web/validation"
	"github.com/pimentafm/auction-go/internal/usecase/bid_usecase"
)

type BidController struct {
	BidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(BidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		BidUseCase: BidUseCase,
	}
}

func (b *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO
	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := b.BidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
