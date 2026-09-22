package main

import (
	"log"
	"time"

	"adamatti.github.io/go-shop/service"
	"github.com/twmb/franz-go/pkg/kgo"
)

var startTime = time.Now()

func main() {
	go service.StartKafkaConsumer(func(record *kgo.Record) {
		log.Printf("received message from kafka topic %s with key %s and value %s", record.Topic, string(record.Key), string(record.Value))
	})
	go startGrpc()
	startHttp()
}
