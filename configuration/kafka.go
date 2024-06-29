package configuration

import "github.com/thankala/gregor_chair_common/enums"

var (
	defaultBrokers   = []string{"localhost:9092"}
	defaultTopic     = enums.AssemblyTask1.String()
	defaultGroupId   = enums.AssemblyTask1.String()
	defaultPartition = 0
	defaultMinBytes  = 10e3 // 10KB
	defaultMaxBytes  = 10e6 // 10MB
)

type KafkaOpts struct {
	Brokers   []string
	Topic     string
	GroupId   string
	Partition int
	MinBytes  int
	MaxBytes  int
}

func DefaultKafkaOpts() *KafkaOpts {
	return &KafkaOpts{
		Brokers:   defaultBrokers,
		Topic:     defaultTopic,
		GroupId:   defaultGroupId,
		Partition: defaultPartition,
		MinBytes:  int(defaultMinBytes),
		MaxBytes:  int(defaultMaxBytes),
	}
}

type KafkaOptionFunc func(*KafkaOpts)

func WithKafkaListenAddresses(brokers ...string) KafkaOptionFunc {
	return func(opts *KafkaOpts) {
		opts.Brokers = brokers
	}
}

func WithKafkaTopic(topics string) KafkaOptionFunc {
	return func(opts *KafkaOpts) {
		opts.Topic = topics
	}
}

func WithKafkaGroupId(groupId string) KafkaOptionFunc {
	return func(opts *KafkaOpts) {
		opts.GroupId = groupId
	}
}

func WithKafkaPartition(partition int) KafkaOptionFunc {
	return func(opts *KafkaOpts) {
		opts.Partition = partition
	}
}

func WithKafkaMinBytes(minBytes int) KafkaOptionFunc {
	return func(opts *KafkaOpts) {
		opts.MinBytes = minBytes
	}
}

func WithKafkaMaxBytes(maxBytes int) KafkaOptionFunc {
	return func(opts *KafkaOpts) {
		opts.MaxBytes = maxBytes
	}
}
