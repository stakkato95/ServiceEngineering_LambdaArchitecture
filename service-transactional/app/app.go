package app

import (
	"github.com/stakkato95/lambda-architecture/service-transactional/domain"
	"github.com/stakkato95/lambda-architecture/service-transactional/logger"
)

func Start() {
	sink := domain.NewUserSink()
	processor := domain.NewUserProcessor(sink)

	if err := processor.StartProcessing(); err != nil {
		logger.Error("can not start processing: " + err.Error())
	}
}
