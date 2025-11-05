package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-spelunker-httpd/app/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}
}
