package metrics

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	// RequestCount Метрика для сбора количества запросов.
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",           // Название метрики
			Help: "Total number of gRPC requests", // Описание метрики
		},
		[]string{"method"},
	)

	// RequestDuration Метрика для сбора времени выполнения запросов в секундах.
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds", // Название метрики
			Help:    "RPC latency distributions",     // Описание метрики
			Buckets: prometheus.DefBuckets,           // Диапазон времени выполнения запросов
		},
		[]string{"method"},
	)

	// ErrorCount Метрика для сбора количества ошибок.
	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_errors_total",                         // Название метрики
			Help: "Total number of gRPC requests with errors", // Описание метрики
		},
		[]string{"method"},
	)
)

const ReadHeaderTimeoutDuration = 5 * time.Second

type Metrics struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Metrics {
	return &Metrics{
		logger: logger,
	}
}

func (m *Metrics) Init() error {
	if err := prometheus.Register(RequestCount); err != nil {
		return fmt.Errorf("prometheus.Register: failed to register RequestCount: %w", err)
	}

	if err := prometheus.Register(RequestDuration); err != nil {
		return fmt.Errorf("prometheus.Register: failed to register RequestDuration: %w", err)
	}

	if err := prometheus.Register(ErrorCount); err != nil {
		return fmt.Errorf("prometheus.Register: failed to register ErrorCount: %w", err)
	}

	return nil
}

func (m *Metrics) Run(ctx context.Context, addr string) error {
	http.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: ReadHeaderTimeoutDuration,
	}

	go func() {
		m.logger.Info("metrics.Run: metrics server started", zap.String("address", addr))

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			m.logger.Error("server.ListenAndServe: ", zap.Error(err))
		}
	}()

	/// Блок остановки сервера по сигналу отмены контекста.
	<-ctx.Done()

	m.logger.Info("metrics.Run: context canceled")

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server.Shutdown: %w", err)
	}

	m.logger.Info("metrics.Run: metrics server shutdown")

	return nil
}
