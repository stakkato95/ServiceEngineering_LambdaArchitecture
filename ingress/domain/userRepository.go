package domain

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stakkato95/lambda-architecture/ingress/config"
	"github.com/stakkato95/lambda-architecture/ingress/errs"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
)

type UserRepository interface {
	InjestUser(User) *errs.AppError
	Destroy() error
}

const userTopic = "user"
const partition = 0

type KafkaUserRepository struct {
	conn *kafka.Conn
}

func NewKafkaUserRepository() UserRepository {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.AppConfig.KafkaService, userTopic, partition)
	if err != nil {
		logger.Fatal("failed to dial leader: " + err.Error())
	} else {
		logger.Info("LEADER WORKS!!")
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return &KafkaUserRepository{conn}
}

func (k *KafkaUserRepository) InjestUser(user User) *errs.AppError {
	_, err := k.conn.WriteMessages(
		kafka.Message{Value: []byte("{ \"name\": \"" + user.Name + "\" }")},
	)
	if err != nil {
		logger.Error("failed to write messages: " + err.Error())
	}
	return nil
}

func (k *KafkaUserRepository) Destroy() error {
	if err := k.conn.Close(); err != nil {
		logger.Fatal("failed to close writer: " + err.Error())
	}

	return nil
}
