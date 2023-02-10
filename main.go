package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafka2 "github.com/falchizao/simulator/application/kafka"
	"github.com/falchizao/simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error while loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.ProduceMSG(msg)
	}

	producer := kafka.NewKafkaProducer(msgChan)
	kafka.Publish("hey", "readtest", producer)

	for {
		_ = 1
	}
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPos()
	// stringjson, _ := route.ExportJSONPos()
	// fmt.Println(stringjson[0])
}
