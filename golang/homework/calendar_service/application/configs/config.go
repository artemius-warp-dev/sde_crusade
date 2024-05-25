package configs

type Storage struct {
	Type    string
	ConnStr *string
}

type Queue struct {
	Type      string
	ConnStr   *string
	QueueName string
}

type ServerConfig struct {
	HTTPAddress string
	GRPCAddress string
	Storage     Storage
	Queue       Queue
}

type HTTPListenConfig struct {
	Port string `yaml:"port"`
	IP   string `yaml:"ip"`
}

type GRPCListenConfig struct {
	Port string `yaml:"port"`
	IP   string `yaml:"ip"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

type RabbitMQConfig struct {
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	QueueName string `yaml:"queue_name"`
}

type Config struct {
	LogLevel   string           `yaml:"log_level"`
	LogFile    string           `yaml:"log_file"`
	HTTPListen HTTPListenConfig `yaml:"http_listen"`
	GRPCListen GRPCListenConfig `yaml:"grpc_listen"`
	Database   DatabaseConfig   `yaml:"database"`
	RabbitMQ   RabbitMQConfig   `yaml:"rabbit"`
}
