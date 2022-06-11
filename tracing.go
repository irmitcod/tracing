package tracing

import (
	"github.com/irmitcod/tracing/config"
	"github.com/irmitcod/tracing/pkg/log"
	"github.com/irmitcod/tracing/pkg/metric"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type tracing struct {
	log.Factory
	metric.Metrics
}

func NewTracing(config config.Config) (*tracing, error) {

	metrics, err := metric.CreateMetrics(config.Metrics.URL, config.Metrics.ServiceName)
	if err != nil {
		return nil, err
	}
	loggerZ, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := loggerZ.With(zap.String("service", "customer"))
	factory := log.NewFactory(zapLogger)

	return &tracing{Factory: factory, Metrics: metrics}, nil
}
