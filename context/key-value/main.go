package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	if token == nil {
		fmt.Println("Token not found")
		return
	}
	fmt.Printf("Token found: %s\n", token)
}
