package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/database/mongodb"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/api/web/controller/auction_controller"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/api/web/controller/bid_controller"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/api/web/controller/user_controller"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/database/auction"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/database/bid"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/infra/database/user"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/auction_usecase"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/bid_usecase"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/user_usecase"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
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
	router.GET("/auctions/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auctions", auctionsController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionsController.FindWinningBidByAuction)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(dataBase *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController,
) {
	auctionRepository := auction.NewAuctionRepository(dataBase)

	bidRepository := bid.NewBidRepository(dataBase, auctionRepository)
	bidUseCase := bid_usecase.NewBidUseCase(bidRepository)

	auctionUseCase := auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository)

	userRepository := user.NewUserRepository(dataBase)
	userUseCase := user_usecase.NewUserUseCase(userRepository)

	userController = user_controller.NewUserController(userUseCase)
	auctionController = auction_controller.NewAuctionController(auctionUseCase)
	bidController = bid_controller.NewBidController(bidUseCase)

	return
}
