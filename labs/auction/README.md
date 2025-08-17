Here is an English README for your project, following the structure and details of your example:

# Golang Concurrency Challenge - Auction

This repository contains the solution for the **Golang Concurrency Challenge - Auction**, part of the **Go Expert / FullCycle** postgraduate program.

## Project Overview

**Goal:**  
Add a new feature to the existing auction project so that auctions automatically close after a defined period.

Clone the repository

The auction and bid creation routines are already implemented. However, the cloned project needs improvement: you must add an automatic closing routine based on time.

For this task, you will use Go routines and focus on the auction creation process. The validation for whether an auction is open or closed during new bids is already implemented.

You must develop:

- A function to calculate the auction duration, based on parameters defined in environment variables.
- A new Go routine that checks for expired auctions and updates their status to closed.
- A test to validate that the closing is happening automatically.

**Tips:**

- Focus on the file `internal/infra/database/auction/create_auction.go` for your implementation.
- Remember to handle concurrency properly.
- Check how the interval calculation for auction validity is performed in the bid creation routine.
- For more information about goroutines, refer to the Multithreading module in the Go Expert course.

## Challenge Details

The main challenge is to implement an automated process that closes auctions after a configurable time interval. This involves:

- Reading the auction interval from environment variables.
- Launching a goroutine during auction creation that waits for the interval and then updates the auction status to closed.
- Ensuring thread safety and proper error handling.
- Writing tests to confirm that auctions are closed automatically after the interval.

## How to Run the Project

- In the project root, run `make`
- To test, run `go test -v internal/infra/database/auction/create_auction_test.go` or `make test`

## Delivery

- Complete source code with the implemented solution.
- Documentation explaining how to run the project in a development environment.
- Use docker/docker-compose so your application can be easily tested.