package domain

import (
	"context"
	"fmt"
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
const msgBufferSize = 10e3 //10KB

type kafkaUserRepository struct {
	conn      *kafka.Conn
	userCount int
}

func NewUserRepository() UserRepository {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.AppConfig.KafkaService, userTopic, partition)
	if err != nil {
		logger.Fatal("failed to dial leader: " + err.Error())
	}

	repo := kafkaUserRepository{conn: conn}

	go func() {
		for {
			//zwischen timeout und anderen Fehlern unterscheiden
			conn.SetReadDeadline(time.Now().Add(60 * time.Second))
			if msg, err := conn.ReadMessage(msgBufferSize); err != nil {
				logger.Error("error when reading a msg from kafka: " + err.Error())
			} else {
				logger.Info(fmt.Sprintf("topic: %s, offset: %d, value: %s", msg.Topic, msg.Offset, string(msg.Value[:])))
				repo.userCount += 1
			}
		}
	}()

	return &repo
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
