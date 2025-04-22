package auction

import (
	"context"
	"time"

	"github.com/pimentafm/auction-go/configuration/logger"
	"github.com/pimentafm/auction-go/internal/entity/auction_entity"
	"github.com/pimentafm/auction-go/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auction struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   time.Time                       `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := Auction{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp,
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to create auction", err)
		return internal_error.NewInternalServerError("Error trying to create auction")
	}
	return nil

}
