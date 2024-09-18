package output

import (
	"context"
	infraPubSub "go-stream/internal/infrastructure/pubsub"
	"go-stream/internal/ports"
	"log"

	gcloudPubSub "cloud.google.com/go/pubsub"
)

type PubSubProducer struct {
	Client *infraPubSub.PubSubClient
}

func NewPubSubProducer(client *infraPubSub.PubSubClient) ports.MessageProducer {
	return &PubSubProducer{Client: client}
}

func (pp *PubSubProducer) Produce(in <-chan interface{}) error {
	ctx := context.Background()
	for data := range in {
		msgData, ok := data.([]byte)
		if !ok {
			log.Printf("Invalid data type: %T", data)
			continue
		}
		result := pp.Client.Topic.Publish(ctx, &gcloudPubSub.Message{
			Data: msgData,
		})
		_, err := result.Get(ctx)
		if err != nil {
			log.Printf("Error publishing message: %v", err)
			return err
		}
	}
	return nil
}
