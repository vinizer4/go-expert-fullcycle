package auction_usecase

import (
	"context"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/auction_entity"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/internal_error"
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
