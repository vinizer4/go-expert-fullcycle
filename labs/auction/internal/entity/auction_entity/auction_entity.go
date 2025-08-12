package auction_entity

import (
	"context"
	"github.com/google/uuid"
	"time"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/internal_error"
)

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductionCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductionCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductionCondition = iota
	Used
	Refurbished
)

type AuctionRepositoryInterface interface {
	FindAuctionById(
		ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAllAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]Auction, *internal_error.InternalError)
	CreateAuction(
		ctx context.Context,
		auctionEntity Auction) *internal_error.InternalError
}

func CreateAuction(
	productName, category, description string,
	condition ProductionCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (au *Auction) Validate() *internal_error.InternalError {
	if len(au.ProductName) <= 1 ||
		len(au.Category) <= 2 ||
		len(au.Description) <= 10 &&
			(au.Condition < New || au.Condition != Refurbished && au.Condition != Used) {
		return internal_error.NewBadRequestError("Invalid auction object")
	}

	return nil
}
