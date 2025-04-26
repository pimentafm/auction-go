package auction_usecase

import (
	"context"
	"time"

	"github.com/pimentafm/auction-go/internal/entity/auction_entity"
	"github.com/pimentafm/auction-go/internal/internal_error"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductCondition(auctionInput.Condition),
	)

	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}

// func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {

// }

// func (au *AuctionUseCase) FindAuctions(
// 	ctx context.Context,
// 	status auction_entity.AuctionStatus,
// 	category, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {

// }
