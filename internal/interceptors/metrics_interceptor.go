package interceptors

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"link_service/internal/server/metrics"
)

// MetricsUnaryInterceptor перехватывает выполнение метода и собирает метрикиере
func MetricsUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		method := info.FullMethod // вытаскивает имя метода который вызывается

		// Увеличиваем значение счетчика вызова указанного метода на 1
		metrics.RequestCount.WithLabelValues(method).Inc()

		// Непосредственно выполняем вызванный метод внутри нашего интерцептора
		resp, err := handler(ctx, req)

		// Считаем длительность запроса
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.WithLabelValues(method).Observe(duration)

		// Если запрос завершился с ошибкой, увеличиваем значение счетчика ошибок на 1
		if err != nil {
			metrics.ErrorCount.WithLabelValues(method).Inc()
		}

		// Возвращаем гРПС ответ и ошибку перехваченного метода
		return resp, err
	}
}
