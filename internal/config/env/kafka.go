package env

import (
	"errors"
	"os"
	"strconv"
)

type kafkaConfig struct {
	hosts string

	userTopic string

	maxRetryCount int
}

func NewKafkaConfig() (*kafkaConfig, error) {
	hosts := os.Getenv("KAFKA_HOSTS")
	if len(hosts) == 0 {
		return nil, errors.New("KAFKA_HOSTS is empty")
	}

	userTopic := os.Getenv("KAFKA_USER_TOPIC")
	if len(userTopic) == 0 {
		return nil, errors.New("KAFKA_USER_TOPIC is empty")
	}

	maxRetryCountStr := os.Getenv("KAFKA_MAX_RETRY_COUNT")
	if len(maxRetryCountStr) == 0 {
		return nil, errors.New("KAFKA_MAX_RETRY_COUNT is empty")
	}

	maxRetryCount, err := strconv.Atoi(maxRetryCountStr)
	if err != nil {
		return nil, errors.New("KAFKA_MAX_RETRY_COUNT is not a number")
	}

	return &kafkaConfig{
		hosts:         hosts,
		userTopic:     userTopic,
		maxRetryCount: maxRetryCount,
	}, nil
}

func (c *kafkaConfig) Hosts() string {
	return c.hosts
}

func (c *kafkaConfig) UserTopic() string {
	return c.userTopic
}

func (c *kafkaConfig) MaxRetryCount() int {
	return c.maxRetryCount
}
