package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/txsvc/stdlib/v2/env"

	"github.com/redhat-partner-ecosystem/openshift-skunkworks/internal"
)

var (
	kc *kafka.Consumer
	//kp *kafka.Producer
)

func init() {

	clientID := env.GetString("client_id", "case-svc")
	groupID := env.GetString("group_id", "fsi-fraud-detection")

	// kafka setup
	kafkaService := env.GetString("kafka_service", "")
	if kafkaService == "" {
		panic(fmt.Errorf("missing env 'kafka_service'"))
	}
	kafkaServicePort := env.GetString("kafka_service_port", "9092")
	kafkaServer := fmt.Sprintf("%s:%s", kafkaService, kafkaServicePort)

	// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
	_kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       kafkaServer,
		"client.id":               clientID,
		"group.id":                groupID,
		"connections.max.idle.ms": 0,
		"auto.offset.reset":       "earliest",
		"broker.address.family":   "v4",
	})
	if err != nil {
		panic(err)
	}
	kc = _kc

	/*
		_kp, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": kafkaServer,
		})
		if err != nil {
			panic(err)
		}
		kp = _kp
	*/

	// prometheus endpoint setup
	internal.StartPrometheusListener()
}

func main() {

	clientID := env.GetString("client_id", "case-svc")
	sourceTopic := env.GetString("source_topic", "tx-fraud")

	// metrics collectors
	opsTxProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "kafka_listener_events",
		Help: "The number of processed events",
	})

	// create a responder for delivery notifications
	evts := make(chan kafka.Event, 1000) // FIXME not sure if such a number is needed ...
	go func() {
		fmt.Printf(" --> %s: listening for delivery events\n", clientID)
		for {
			e := <-evts

			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf(" --> delivery error: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// subscriber
	err := kc.SubscribeTopics([]string{sourceTopic}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf(" --> %s: listening on topic '%s'\n", clientID, sourceTopic)

	for {
		msg, err := kc.ReadMessage(-1)
		if err == nil {
			fmt.Printf(" ---> message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			/*
				// back to a json string
				data, err := json.Marshal(tx)
				if err != nil {
					// do something
				}

				// send to the next destination
				err = kp.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic:     &targetTopic,
						Partition: kafka.PartitionAny,
					},
					Value: data,
				}, evts)
				if err != nil {
					fmt.Printf(" --> producer error: %v\n", err)
				}
			*/

			// metrics
			opsTxProcessed.Inc()

		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf(" --> consumer error: %v (%v)\n", err, msg)
		}
	}
}
