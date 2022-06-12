package jaeger

import (
	"github.com/irmitcod/tracing/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

const id = 1

// Init Jaeger
func InitJaeger(cfg *config.Config) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.Jaeger.Host)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Jaeger.ServiceName),
			attribute.String("environment", "production"),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
	//jaegerCfgInstance := jaegercfg.Configuration{
	//	ServiceName: cfg.Jaeger.ServiceName,
	//	Sampler: &jaegercfg.SamplerConfig{
	//		Type:  jaeger.SamplerTypeConst,
	//		Param: 1,
	//	},
	//	Reporter: &jaegercfg.ReporterConfig{
	//		LogSpans:           cfg.Jaeger.LogSpans,
	//		LocalAgentHostPort: cfg.Jaeger.Host,
	//	},
	//}
	//
	//return jaegerCfgInstance.NewTracer(
	//	jaegercfg.Logger(jaegerlog.StdLogger),
	//	jaegercfg.Metrics(metrics.NullFactory),
	//)
}
