package domain

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/stakkato95/lambda-architecture/processor-analytics/config"
	"github.com/stakkato95/lambda-architecture/processor-analytics/logger"
)

type UserSink interface {
	Sink(int) error
}

type cassandraUserSink struct {
	session *gocql.Session
}

func NewUserSink() UserSink {
	cluster := gocql.NewCluster(config.AppConfig.CassandraCluster)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.AppConfig.CassandraUser,
		Password: config.AppConfig.CassandraPassword,
	}
	cluster.Keyspace = config.AppConfig.CassandraKeyspace
	cluster.Timeout = 10 * time.Second

	cluster.ProtoVersion = 4
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Fatal("Could not connect to cassandra cluster: " + err.Error())
	}

	return &cassandraUserSink{session}
}

func (s *cassandraUserSink) Sink(userCount int) error {
	logger.Info(fmt.Sprintf("sink user: %d", userCount))

	uuid, err := gocql.RandomUUID()
	if err != nil {
		logger.Fatal("can not generate random uuid: " + uuid.String())
	}

	err = s.session.Query(
		"INSERT INTO "+config.AppConfig.CassandraTable+" (id, time, user_count) VALUES (?, ?, ?)",
		uuid.String(), time.Now(), userCount).Exec()
	if err != nil {
		logger.Error("can not insert user count: " + err.Error())
	}

	return nil
}
