package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/stakkato95/lambda-architecture/service-transactional/config"
	"github.com/stakkato95/lambda-architecture/service-transactional/logger"
)

type UserProcessor interface {
	StartProcessing() error
}

const partition = 0
const msgBufferSize = 10e3 //10KB

type kafkaUserProcessor struct {
	conn *kafka.Conn
	sink UserSink
}

func NewUserProcessor(sink UserSink) UserProcessor {
	return &kafkaUserProcessor{sink: sink}
}

func (p *kafkaUserProcessor) StartProcessing() error {
	var err error
	kafkaService := config.AppConfig.KafkaService
	topic := config.AppConfig.KafkaTopic
	if p.conn, err = kafka.DialLeader(context.Background(), "tcp", kafkaService, topic, partition); err != nil {
		return err
	}

	for {
		//zwischen timeout und anderen Fehlern unterscheiden
		p.conn.SetReadDeadline(time.Now().Add(60 * time.Hour))
		if msg, err := p.conn.ReadMessage(msgBufferSize); err != nil {
			logger.Error("error when reading a msg from kafka: " + err.Error())
		} else {
			var user User
			json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&user)
			logger.Info(fmt.Sprintf("read user: %v", user))
			p.sink.Sink(user)
		}
	}
}
