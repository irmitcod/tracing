package main

import (
	"context"
	"github.com/irmitcod/tracing/config"
	"github.com/irmitcod/tracing/pkg/jaeger"
	"github.com/irmitcod/tracing/pkg/log"
	"github.com/irmitcod/tracing/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func main() {

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)

	loggerZ, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := loggerZ.With(zap.String("service", "customer"))
	factoryLogger := log.NewFactory(zapLogger)
	//factory.

	tp, err := jaeger.InitJaeger(cfg)
	if err != nil {
		factoryLogger.Bg().Fatal("cannot create tracer", zap.Error(err))
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	factoryLogger.Bg().Info("Jaeger connected")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			factoryLogger.Bg().Fatal(err.Error())
		}
	}(ctx)

}
