package input

import (
	"context"
	infraPubSub "go-stream/internal/infrastructure/pubsub"
	"go-stream/internal/ports"
	"log"
	"sync"

	gcloudPubSub "cloud.google.com/go/pubsub"
)

type PubSubConsumer struct {
	Client *infraPubSub.PubSubClient
}

func NewPubSubConsumer(client *infraPubSub.PubSubClient) ports.MessageConsumer {
	return &PubSubConsumer{Client: client}
}

func (pc *PubSubConsumer) Consume(out chan<- interface{}) error {
	ctx := context.Background()
	var wg sync.WaitGroup

	err := pc.Client.Subscription.Receive(ctx, func(ctx context.Context, msg *gcloudPubSub.Message) {
		wg.Add(1)
		defer wg.Done()
		out <- msg.Data
		msg.Ack()
	})

	wg.Wait()

	if err != nil {
		log.Printf("Error receiving messages: %v", err)
		return err
	}

	return nil
}
