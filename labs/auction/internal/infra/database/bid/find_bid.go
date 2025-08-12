package bid

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/logger"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/bid_entity"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/internal_error"
)

func (bd *BidRepository) FindBidByAuctionId(
	ctx context.Context,
	auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auction_id": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)

	if err != nil {
		logger.Error(
			fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId))
	}

	var bidEntitiesMongo []BidEntityMongo
	if err = cursor.All(ctx, &bidEntitiesMongo); err != nil {
		logger.Error(
			fmt.Sprintf("Error trying to decode bids by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf("Error trying to decode bids by auctionId %s", auctionId))
	}

	var bidEntities []bid_entity.Bid
	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(
	ctx context.Context,
	auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auction_id": auctionId}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error("Error trying to find the winning bid by auctionId", err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf("Error trying to find the winni¯ng bid by auctionId %s", auctionId))
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
