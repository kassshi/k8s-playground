package middleware

import (
	"context"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	"go.opentelemetry.io/otel/trace"
)

func NewAccessLoggingInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()
			method := req.Spec().Procedure
			res, err := next(ctx, req)
			logger := slog.Default()
			spanContext := trace.SpanContextFromContext(ctx)
			if err != nil {
				logger.ErrorContext(ctx, "rpc request completed", "rpc.procedure", method, "rpc.code", connect.CodeOf(err), "duration", time.Since(start).String(), "trace_id", spanContext.TraceID().String(), "span_id", spanContext.SpanID().String(), "err", err.Error())
			} else {
				logger.InfoContext(ctx, "rpc request completed", "rpc.procedure", method, "rpc.code", "ok", "duration", time.Since(start).String(), "trace_id", spanContext.TraceID().String(), "span_id", spanContext.SpanID().String())
			}
			return res, err
		}
	}
}
