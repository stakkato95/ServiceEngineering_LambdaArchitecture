package domain

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/stakkato95/lambda-architecture/service-analytics/config"
	"github.com/stakkato95/lambda-architecture/service-analytics/logger"
)

type UserCountRepository interface {
	GetUserCounts() []UserCount
}

type cassandraUserCountRepository struct {
	session *gocql.Session
}

func NewUserCountRepository() UserCountRepository {
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

	return &cassandraUserCountRepository{session}
}

func (r *cassandraUserCountRepository) GetUserCounts() []UserCount {
	iter := r.session.Query("SELECT id, time, user_count FROM user").Iter()
	counts := []UserCount{}

	for {
		var userCount UserCount
		if !iter.Scan(&userCount.Id, &userCount.Time, &userCount.UserCount) {
			break
		}
		counts = append(counts, userCount)
	}

	return counts
}
