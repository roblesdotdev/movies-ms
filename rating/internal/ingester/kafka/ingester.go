package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/roblesdotdev/movies-ms/rating/pkg/model"
)

type Ingester struct {
	consumer *kafka.Consumer
	topic    string
}

func NewIngester(addr string, groupId string, topic string) (*Ingester, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &Ingester{consumer: consumer, topic: topic}, nil
}

func (i *Ingester) Ingest(ctx context.Context) (chan model.RatingEvent, error) {
	if err := i.consumer.SubscribeTopics([]string{i.topic}, nil); err != nil {
		return nil, err
	}
	ch := make(chan model.RatingEvent, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				i.consumer.Close()
				return
			default:
			}
			msg, err := i.consumer.ReadMessage(-1)
			if err != nil {
				fmt.Println("consumer error: " + err.Error())
				continue
			}
			var event model.RatingEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				fmt.Println("unmarshal error: " + err.Error())
				continue
			}
			ch <- event
		}
	}()
	return ch, nil
}
