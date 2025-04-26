package auction_controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pimentafm/auction-go/configuration/rest_err"
	"github.com/pimentafm/auction-go/internal/usecase/auction_usecase"
)

func (a *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("auctionId")
	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID format",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := a.AuctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

func (a *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	prodductName := c.Query("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("Error traying to validate auction status param")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := a.AuctionUseCase.FindAuctions(context.Background(), auction_usecase.AuctionStatus(statusNumber), category, prodductName)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)
}

func (a *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")
	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID format",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := a.AuctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}
