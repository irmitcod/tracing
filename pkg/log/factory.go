// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Factory is the default logging wrapper that can create
// logger instances either for a given Context or context-less.
type Factory struct {
	logger *zap.Logger
}

// NewFactory creates a new Factory.
func NewFactory(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

// Bg creates a context-unaware logger.
func (b Factory) Bg() Logger {
	return logger(b)
}

func (b Factory) GetGloabalCntext() trace.TracerProvider {
	return otel.GetTracerProvider()
}

// ForError returns a context-aware Logger. If the context
// contains an OpenTracing span, all logging calls are also
// echo-ed into the span.
func (b Factory) ForError(ctx context.Context) Logger {
	tr := otel.Tracer("ForError")
	_, span := tr.Start(ctx, "ForInfo")

	logger := spanLogger{span: span, logger: b.logger}
	logger.span.SetStatus(1, "error")
	jaegerCtx := span.SpanContext()
	logger.spanFields = []zapcore.Field{
		zap.String("trace_id", jaegerCtx.TraceID().String()),
		zap.String("span_id", jaegerCtx.SpanID().String()),
	}

	return logger

}

// ForInfo returns a context-aware Logger. If the context
// contains an OpenTracing span, all logging calls are also
// echo-ed into the span.
func (b Factory) ForInfo(ctx context.Context) Logger {

	tr := otel.Tracer("ForInfo")
	_, span := tr.Start(ctx, "ForInfo")
	logger := spanLogger{span: span, logger: b.logger}

	jaegerCtx := span.SpanContext()
	logger.spanFields = []zapcore.Field{
		zap.String("trace_id", jaegerCtx.TraceID().String()),
		zap.String("span_id", jaegerCtx.SpanID().String()),
	}

	return logger

}

// With creates a child logger, and optionally adds some context fields to that logger.
func (b Factory) With(fields ...zapcore.Field) Factory {
	return Factory{logger: b.logger.With(fields...)}
}
