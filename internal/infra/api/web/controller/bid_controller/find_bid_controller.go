package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pimentafm/auction-go/configuration/rest_err"
)

func (a *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")
	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID format",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	bidOutputList, err := a.BidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bidOutputList)
}
