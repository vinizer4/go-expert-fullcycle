package auction

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
	"time"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/database/mongodb"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/auction_entity"
)

func TestCreateAuction_ExpireStatus(t *testing.T) {
	ctx := context.Background()

	os.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	os.Setenv("MONGODB_DB", "auction_test")
	os.Setenv("AUCTION_INTERVAL", "2s")

	db, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		t.Fatalf("MongoDB connection error: %+v", err)
	}

	ar := NewAuctionRepository(db)

	a := auction_entity.Auction{
		Id:          "test-auction-id-001",
		ProductName: "test_product",
		Category:    "test_category",
		Description: "test_description",
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	if err := ar.CreateAuction(ctx, a); err != nil {
		t.Fatalf("failed to create auction: %+v | type: %T", err, err)
	}

	var result AuctionEntityMongo
	err = ar.Collection.FindOne(ctx, bson.M{"_id": a.Id}).Decode(&result)
	if err != nil {
		t.Fatalf("auction not found after creation: %+v", err)
	}
	if result.Status != auction_entity.Active {
		t.Errorf("auction status is not active after creation")
	}

	time.Sleep(3 * time.Second)

	err = ar.Collection.FindOne(ctx, bson.M{"_id": a.Id}).Decode(&result)
	if err != nil {
		t.Fatalf("auction not found after interval: %+v", err)
	}
	if result.Status != auction_entity.Completed {
		t.Errorf("auction status was not updated to completed after interval")
	}

	_, _ = ar.Collection.DeleteOne(ctx, bson.M{"_id": a.Id})
}
