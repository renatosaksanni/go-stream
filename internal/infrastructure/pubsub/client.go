package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type PubSubClient struct {
	Client       *pubsub.Client
	Subscription *pubsub.Subscription
	Topic        *pubsub.Topic
}

func NewPubSubClient(projectID, subscriptionID, topicID string) (*PubSubClient, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	sub := client.Subscription(subscriptionID)
	topic := client.Topic(topicID)

	return &PubSubClient{
		Client:       client,
		Subscription: sub,
		Topic:        topic,
	}, nil
}
