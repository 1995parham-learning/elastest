package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/1995parham-learning/elastest/actions"
)

const ShutdownTimeout = 5 * time.Second

func main() {
	log.Println("18.20 at Sep 07 2016 7:20 IR721")

	app := actions.App()

	go func() {
		if err := app.Start(":8080"); err != http.ErrServerClosed {
			log.Fatalf("API Service failed with %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("18.20 As always ... left me alone")

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Printf("API Service failed on exit: %s", err)
	}
}
