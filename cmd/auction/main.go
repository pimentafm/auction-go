package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pimentafm/auction-go/configuration/database/mongodb"
	"github.com/pimentafm/auction-go/internal/infra/api/web/controller/auction_controller"
	"github.com/pimentafm/auction-go/internal/infra/api/web/controller/bid_controller"
	"github.com/pimentafm/auction-go/internal/infra/api/web/controller/user_controller"
	"github.com/pimentafm/auction-go/internal/infra/database/auction"
	"github.com/pimentafm/auction-go/internal/infra/database/bid"
	"github.com/pimentafm/auction-go/internal/infra/database/user"
	"github.com/pimentafm/auction-go/internal/usecase/auction_usecase"
	"github.com/pimentafm/auction-go/internal/usecase/bid_usecase"
	"github.com/pimentafm/auction-go/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionsController := initDependencies(databaseConnection)

	router.GET("/auctions", auctionsController.FindAuctions)
	router.GET("/auction/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auctions", auctionsController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController,
) {
	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))

	return
}
