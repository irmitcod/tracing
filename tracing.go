package tracing

import (
	"context"
	"github.com/irmitcod/tracing/config"
	"github.com/irmitcod/tracing/pkg/jaeger"
	"github.com/irmitcod/tracing/pkg/log"
	"github.com/irmitcod/tracing/pkg/metric"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type tracing struct {
	log.Factory
	metric.Metrics
}

func NewTracing(config *config.Config) (*tracing, error) {

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

	tp, err := jaeger.InitJaeger(config)
	if err != nil {
		return nil, err
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			zapLogger.Fatal(err.Error())
		}
	}(ctx)


	return &tracing{Factory: factory, Metrics: metrics}, nil
}
