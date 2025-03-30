package config

import (
	"github.com/poteto-go/poteto"
	"github.com/poteto-go/poteto/middleware"
	"github.com/sirupsen/logrus"
)

func NewRequestLoggerConfig() middleware.RequestLoggerConfig {
	log := logrus.New()
	logConfig := middleware.DefaultRequestLoggerConfig

	logConfig.LogHandleFunc = func(ctx poteto.Context, rlv middleware.RequestLoggerValues) error {
		if rlv.Error == nil {
			log.WithFields(logrus.Fields{
				"method":    rlv.Method,
				"routePath": rlv.RoutePath,
				"status":    rlv.Status,
				"realIP":    rlv.RealIP,
				"requestId": rlv.RequestId,
			}).Info("request")
		} else {
			log.WithFields(logrus.Fields{
				"method":    rlv.Method,
				"routePath": rlv.RoutePath,
				"status":    rlv.Status,
				"realIP":    rlv.RealIP,
				"requestId": rlv.RequestId,
			}).Error("request")
		}
		return nil
	}

	return logConfig
}
