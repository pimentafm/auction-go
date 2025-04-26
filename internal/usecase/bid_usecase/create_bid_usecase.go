package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/pimentafm/auction-go/configuration/logger"
	"github.com/pimentafm/auction-go/internal/entity/bid_entity"
	"github.com/pimentafm/auction-go/internal/internal_error"
)

type BidInputDTO struct {
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidRepository

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

func NewBidUseCase(bidRepository bid_entity.BidRepository) BidUseCaseInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()

	bidUsecase := &BidUseCase{
		BidRepository:       bidRepository,
		timer:               time.NewTimer(maxSizeInterval),
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: getMaxBatchSizeInterval(),
		bidChannel:          make(chan bid_entity.Bid, maxBatchSize),
	}

	bidUsecase.triggerCreateRoutine(context.Background())

	return bidUsecase
}

var bidBatch []bid_entity.Bid

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
}

func (bu *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(bu.bidChannel)

		for {
			select {
			case bidEntity, ok := <-bu.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("Error creating bid batch: ", err)
						}
					}
					return
				}

				bidBatch = append(bidBatch, bidEntity)
				if len(bidBatch) >= bu.maxBatchSize {
					if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("Error creating bid batch: ", err)
					}
					bidBatch = nil
					bu.timer.Reset(bu.batchInsertInterval)
				}
			case <-bu.timer.C:
				if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("Error creating bid batch: ", err)
				}
				bidBatch = nil
				bu.timer.Reset(bu.batchInsertInterval)
			}
		}
	}()
}

func (bu *BidUseCase) CreateBid(
	ctx context.Context,
	bidInputDTO BidInputDTO) *internal_error.InternalError {

	bidEntity, err := bid_entity.CreateBid(bidInputDTO.UserId, bidInputDTO.AuctionId, bidInputDTO.Amount)
	if err != nil {
		return err
	}

	bu.bidChannel <- *bidEntity

	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)

	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getMaxBatchSize() int {
	value, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}

	return value
}
