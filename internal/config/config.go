package config

type Config struct {
	ServiceName    string   `envconfig:"SERVICE_NAME" required:"true"`
	Version        string   `envconfig:"VERSION" required:"true"`
	GRPCPort       string   `envconfig:"GRPC_PORT" default:"50051"`
	HTTPPort       string   `envconfig:"HTTP_PORT" default:"8080"`
	PrometheusPort string   `envconfig:"PROMETHEUS_PORT" default:"9090"`
	LogLevel       string   `envconfig:"LOG_LEVEL" default:"debug"`      // Уровень логирования
	GatewayEnable  bool     `envconfig:"GATEWAY_ENABLE" default:"false"` // Флаг для включения HTTP-сервера
	Postgres       Postgres `envconfig:"POSTGRES"`
	Redis          Redis    `envconfig:"REDIS"`
	Kafka          Kafka    `envconfig:"KAFKA"`
}

type Postgres struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     string `envconfig:"PORT" required:"true"`
	Username string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Database string `envconfig:"DATABASE" required:"true"`
}

type Redis struct {
	Addr     string `envconfig:"ADDR" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DB       int    `envconfig:"DB" required:"true"`
}

type Kafka struct {
	Addr  string `envconfig:"ADDR" default:"9092"`
	Topic string `envconfig:"TOPIC" default:"link_service"`
}
