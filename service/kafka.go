package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

var kafkaClient *kgo.Client

const (
	topic = "orders"
	group = "go-shop"
)

func init() {
	var err error
	kafkaClient, err = kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics(topic),
		kgo.ConsumerGroup(group),
		kgo.DefaultProduceTopic(topic),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.AllowAutoTopicCreation(),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
		kgo.DisableAutoCommit(),
	)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}
}

func SendToKafka(topic string, key []byte, message interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	r := &kgo.Record{
		Topic: topic,
		Key:   key,
		Value: body,
	}

	if err != nil {
		log.Printf("failed to marshal message: %v", err)
		return err
	}

	results := kafkaClient.ProduceSync(ctx, r)
	err = results.FirstErr()

	if err == nil {
		log.Printf("message sent to kafka topic %s with key %s", topic, string(key))
	}

	return err
}

func StartKafkaConsumer(cb func(record *kgo.Record)) {
	ctx := context.Background()
	var fetches kgo.Fetches
	for {
		fetches = kafkaClient.PollFetches(ctx)
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			cb(record)
		}
		if err := kafkaClient.CommitUncommittedOffsets(ctx); err != nil {
			log.Printf("commit: %v", err)
		}
	}
}
