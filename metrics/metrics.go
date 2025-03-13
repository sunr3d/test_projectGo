package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (

	// RequestCount Метрика для сбора количества запросов
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",           // Название метрики
			Help: "Total number of gRPC requests", // Описание метрики
		},
		[]string{"method"},
	)

	// RequestDuration Метрика для сбора времени выполнения запросов в секундах
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds", // Название метрики
			Help:    "RPC latency distributions",     // Описание метрики
			Buckets: prometheus.DefBuckets,           // Диапазон времени выполнения запросов
		},
		[]string{"method"},
	)

	// ErrorCount Метрика для сбора количества ошибок
	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_errors_total",                         // Название метрики
			Help: "Total number of gRPC requests with errors", // Описание метрики
		},
		[]string{"method"},
	)
)

func Init() {
	// Регистрация метрик
	prometheus.MustRegister(RequestCount, RequestDuration, ErrorCount)
}

func NewHandler() http.Handler {
	// Создание обработчика для метрик
	return promhttp.Handler()
}
