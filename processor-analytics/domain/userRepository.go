package domain

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stakkato95/lambda-architecture/processor-analytics/config"
	"github.com/stakkato95/lambda-architecture/processor-analytics/logger"
)

type UserRepository interface {
	GetUserCount() int
	Destroy() error
}

const userTopic = "user"
const partition = 0

type kafkaUserRepository struct {
	conn      *kafka.Conn
	userCount int
}

func NewUserRepository() UserRepository {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.AppConfig.KafkaService, userTopic, partition)
	if err != nil {
		logger.Fatal("failed to dial leader: " + err.Error())
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return &kafkaUserRepository{conn, 100900}
}

func (k *kafkaUserRepository) GetUserCount() int {
	return k.userCount
}

func (k *kafkaUserRepository) Destroy() error {
	if err := k.conn.Close(); err != nil {
		logger.Fatal("failed to close writer: " + err.Error())
	}

	return nil
}
