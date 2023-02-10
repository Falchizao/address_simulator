package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/falchizao/simulator/application/route"
	"github.com/falchizao/simulator/infra/kafka"
)

func ProduceMSG(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	router := route.NewRoute()
	json.Unmarshal(msg.Value, &router)
	router.LoadPos()
	positions, err := route.NewRoute().ExportJSONPos()

	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
