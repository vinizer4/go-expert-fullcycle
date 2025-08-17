package auction_usecase

import (
	"context"
	"time"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/auction_entity"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/bid_entity"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/internal_error"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/bid_usecase"
)

type AuctionInputDto struct {
	ProductName string              `json:"product_name" validate:"required"`
	Category    string              `json:"category" validate:"required"`
	Description string              `json:"description" validate:"required"`
	Condition   ProductionCondition `json:"condition" validate:"required"`
}

type AuctionOutputDto struct {
	Id          string              `json:"id"`
	ProductName string              `json:"product_name"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	Condition   ProductionCondition `json:"condition"`
	Status      AuctionStatus       `json:"status"`
	Timestamp   time.Time           `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDto struct {
	Auction AuctionOutputDto          `json:"auction_id"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type ProductionCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidEntityRepository
}

func (au *AuctionUseCase) CreateAuction(
	ctx context.Context,
	auctionInputDto AuctionInputDto,
) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInputDto.ProductName,
		auctionInputDto.Category,
		auctionInputDto.Description,
		auction_entity.ProductionCondition(auctionInputDto.Condition),
	)

	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, *auction); err != nil {
		return err
	}

	return nil
}
