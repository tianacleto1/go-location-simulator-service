package main

import (
	"fmt"
	"log"

	kafka2 "com.anacleto.location-simulator/application/kafka"
	"com.anacleto.location-simulator/config/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

//{"clientId":"1", "routeId":"1"}
//{"clientId":"2", "routeId":"2"}
//{"clientId":"3", "routeId":"3"}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file!")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
