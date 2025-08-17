package auction_usecase

import (
	"context"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/auction_entity"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/internal_error"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/bid_usecase"
)

func (au *AuctionUseCase) FindAuctionById(
	ctx context.Context,
	id string,
) (*AuctionOutputDto, *internal_error.InternalError) {
	auctionEntity, err := au.auctionRepositoryInterface.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDto{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductionCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil
}

func (au *AuctionUseCase) FindAllAuctions(
	ctx context.Context,
	status AuctionStatus,
	category, productName string,
) ([]AuctionOutputDto, *internal_error.InternalError) {
	auctionEntities, err := au.auctionRepositoryInterface.FindAllAuctions(
		ctx,
		auction_entity.AuctionStatus(status),
		category,
		productName,
	)
	if err != nil {
		return nil, err
	}

	var AuctionOutputDtos []AuctionOutputDto

	for _, value := range auctionEntities {
		AuctionOutputDtos = append(AuctionOutputDtos, AuctionOutputDto{
			Id:          value.Id,
			ProductName: value.ProductName,
			Category:    value.Category,
			Description: value.Description,
			Condition:   ProductionCondition(value.Condition),
			Status:      AuctionStatus(value.Status),
			Timestamp:   value.Timestamp,
		})
	}

	return AuctionOutputDtos, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionId(
	ctx context.Context,
	auctionId string,
) (*WinningInfoOutputDto, *internal_error.InternalError) {
	auction, err := au.auctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionOutputDto := AuctionOutputDto{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductionCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	bidWinning, err := au.bidRepositoryInterface.FindWinningBidByAuctionId(ctx, auction.Id)
	if err != nil {
		return &WinningInfoOutputDto{
			Auction: auctionOutputDto,
			Bid:     nil,
		}, nil
	}

	bidOutputDto := &bid_usecase.BidOutputDTO{
		Id:        bidWinning.Id,
		UserId:    bidWinning.UserId,
		AuctionId: bidWinning.AuctionId,
		Amount:    bidWinning.Amount,
		Timestamp: bidWinning.Timestamp,
	}

	return &WinningInfoOutputDto{
		Auction: auctionOutputDto,
		Bid:     bidOutputDto,
	}, nil
}
