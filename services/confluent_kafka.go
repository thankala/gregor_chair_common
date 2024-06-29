package services

import (
	"encoding/json"
	"github.com/anthdm/hollywood/actor"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/thankala/gregor_chair_common/configuration"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/logger"
	"github.com/thankala/gregor_chair_common/messages"
	"strings"
	"time"
)

type ConfluentKafkaServer struct {
	reader *kafka.Consumer
	writer *kafka.Producer
	stopCh chan struct{}
}

func NewConfluentKafkaServer(opts ...configuration.KafkaOptionFunc) *ConfluentKafkaServer {
	options := configuration.DefaultKafkaOpts()
	for _, opt := range opts {
		opt(options)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(options.Brokers, ","),
		"group.id":                 options.GroupId,
		"auto.offset.reset":        "earliest",
		"enable.auto.commit":       "true",
		"allow.auto.create.topics": "true",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{options.Topic}, nil)

	if err != nil {
		panic(err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(options.Brokers, ","),
		"allow.auto.create.topics": "true",
	})

	if err != nil {
		panic(err)
	}

	return &ConfluentKafkaServer{
		reader: consumer,
		writer: producer,
	}
}

func (k *ConfluentKafkaServer) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		// Do nothing
	case actor.Started:
		go k.Accept(ctx, k.stopCh)
	case actor.Stopped:
		close(k.stopCh)
	case *messages.AssemblyTaskMessage:
		k.Send(msg.Source, msg.Destination, msg.Event, msg)
	case *messages.CoordinatorMessage:
		k.Send(msg.Source, msg.Destination, msg.Event, msg)
	default:
		logger.Get().Error("Unknown message", "Message", msg)
	}
}

func (k *ConfluentKafkaServer) Accept(ctx *actor.Context, stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return // Exit the loop if something is received on stopCh
		default:
			m, err := k.reader.ReadMessage(time.Hour)
			if err != nil {
				logger.Get().Error("Unable to receive events from Kafka", "Error", err)
				continue
			}
			var baseEvent messages.BaseEvent
			err = json.Unmarshal(m.Value, &baseEvent)
			if err != nil {
				logger.Get().Error("Unable to serialize event", "Event", m, "Error", err)
				continue
			}
			switch baseEvent.Event {
			case enums.CoordinatorEvent:
				var coordinatorMessage messages.CoordinatorMessage
				err = json.Unmarshal(baseEvent.Data, &coordinatorMessage)
				if err == nil {
					// Send the message to the parent actor
					ctx.Send(ctx.Parent(), &coordinatorMessage)
					continue
				}
			case enums.AssemblyTaskEvent:
				var assemblyTaskMessage messages.AssemblyTaskMessage
				err = json.Unmarshal(baseEvent.Data, &assemblyTaskMessage)
				if err == nil {
					// Send the message to the parent actor
					ctx.Send(ctx.Parent(), &assemblyTaskMessage)
					continue
				}
			default:
				logger.Get().Warn("Failed to map event", "Event", m)
			}
		}
	}
}

func (k *ConfluentKafkaServer) Send(from string, to string, event enums.Event, message any) {
	// Marshal the message into JSON bytes
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		logger.Get().Error("Unable to marshal event", "Event", message)
		return
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range k.writer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logger.Get().Error("Delivery failed", ev.TopicPartition)
				} else {
					logger.Get().Info("Delivered message", ev.TopicPartition)
				}
			}
		}
	}()

	// Create a BaseEvent struct with the marshaled message as RawMessage
	baseEvent := &messages.BaseEvent{
		Event: event,
		Data:  json.RawMessage(jsonMessage), // Assign marshaled message to Data
	}

	// Marshal the BaseEvent struct which includes the event type and the RawMessage
	m, err := json.Marshal(baseEvent)
	if err != nil {
		logger.Get().Error("Unable to marshal event", "Event", message)
		return
	}
	err = k.writer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &to, Partition: kafka.PartitionAny},
		Key:            []byte(from),
		Value:          m,
	}, nil)
	if err != nil {
		logger.Get().Error("Unable to send event", "Event", message)
	}
}

func (k *ConfluentKafkaServer) GetProducer() actor.Producer {
	return func() actor.Receiver {
		return k
	}
}
