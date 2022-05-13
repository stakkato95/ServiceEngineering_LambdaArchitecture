package app

import (
	"github.com/stakkato95/lambda-architecture/processor-batch/domain"
	"github.com/stakkato95/lambda-architecture/processor-batch/logger"
)

func Start() {
	processor := domain.NewUserProcessor()

	if err := processor.DoProcessing(); err != nil {
		logger.Error("can not perform processing: " + err.Error())
	}
}
