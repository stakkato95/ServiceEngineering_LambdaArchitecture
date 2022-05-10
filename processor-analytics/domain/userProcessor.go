package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stakkato95/lambda-architecture/processor-analytics/config"
	"github.com/stakkato95/lambda-architecture/processor-analytics/logger"
)

type UserProcessor interface {
	GetUserCount() int
	Destroy() error
}

const partition = 0
const msgBufferSize = 10e3 //10KB

type kafkaUserProcessor struct {
	conn      *kafka.Conn
	userCount int
	sink      UserSink
}

func NewUserProcessor(sink UserSink) UserProcessor {
	kafkaService := config.AppConfig.KafkaService
	topic := config.AppConfig.KafkaTopic
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaService, topic, partition)
	if err != nil {
		logger.Fatal("failed to dial leader: " + err.Error())
	}

	repo := kafkaUserProcessor{conn: conn, sink: sink}

	go func() {
		for {
			//zwischen timeout und anderen Fehlern unterscheiden
			conn.SetReadDeadline(time.Now().Add(60 * time.Hour))
			if msg, err := conn.ReadMessage(msgBufferSize); err != nil {
				logger.Error("error when reading a msg from kafka: " + err.Error())
			} else {
				// logger.Info(fmt.Sprintf("topic: %s, offset: %d, value: %s", msg.Topic, msg.Offset, string(msg.Value[:])))
				repo.userCount += 1

				var user User
				json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&user)
				logger.Info(fmt.Sprintf("count: %d, read user: %v", repo.userCount, user))
				repo.sink.Sink(repo.userCount)
			}
		}
	}()

	return &repo
}

func (k *kafkaUserProcessor) GetUserCount() int {
	return k.userCount
}

func (k *kafkaUserProcessor) Destroy() error {
	if err := k.conn.Close(); err != nil {
		logger.Fatal("failed to close writer: " + err.Error())
	}

	return nil
}
