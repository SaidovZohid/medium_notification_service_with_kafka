package main

import (
	"context"
	"log"

	"github.com/SaidovZohid/medium_notification_service_with_kafka/config"
	"github.com/SaidovZohid/medium_notification_service_with_kafka/events"
)

func main() {
	cfg := config.Load(".")

	pubSub, err := events.New(cfg)
	if err != nil {
		log.Fatalf("failed to open pubsub: %v", err)
	}

	pubSub.Run(context.Background())
}
