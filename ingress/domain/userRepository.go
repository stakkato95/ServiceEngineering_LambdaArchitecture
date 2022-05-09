package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

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
	}

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return &KafkaUserRepository{conn}
}

func (k *KafkaUserRepository) InjestUser(user User) *errs.AppError {
	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		logger.Fatal("can not encode user struct: " + err.Error())
	}

	if bytesWritten, err := k.conn.Write(w.Bytes()); err != nil {
		logger.Error("failed to write messages: " + err.Error())
		return errs.NewInjestError(err.Error())
	} else {
		logger.Info(fmt.Sprintf("writter user: %s, written bytes: %d", w.String(), bytesWritten))
		return nil
	}
}

func (k *KafkaUserRepository) Destroy() error {
	if err := k.conn.Close(); err != nil {
		logger.Fatal("failed to close writer: " + err.Error())
	}

	return nil
}
