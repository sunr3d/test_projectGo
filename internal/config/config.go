package config

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" required:"true"`
	Version     string `envconfig:"VERSION" required:"true"`
	GRPCPort    string `envconfig:"GRPC_PORT" default:"50051"`
	//PrometheusPort string `envconfig:"PROMETHEUS_PORT" default:"9090"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"` // Уровень логирования
}
