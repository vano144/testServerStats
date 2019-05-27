package configs

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Config = InitConfig()

// Config struct with read from env file and use in packages
type SConfig struct {
	LogLevel       string `env:"LOG_LEVEL" envDefault:"debug"`
	ServerPort     string `env:"SERVER_PORT" envDefault:":8080"`
	ServiceName    string `env:"SERVICE_NAME" envDefault:"testServerStats"`

	// Db configuration
	DatabaseHost    string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort    int    `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName    string `env:"DATABASE_NAME" envDefault:""`
	DatabaseUser    string `env:"DATABASE_USER" envDefault:""`
	DatabasePass    string `env:"DATABASE_PASS" envDefault:""`

	MaxConnections       int  `env:"MAX_CONNECTIONS" envDefault:"5"`
	MaxIdleConnections   int  `env:"MAX_IDLE_CONNECTIONS" envDefault:"5"`
	MaxOpenConnections   int  `env:"MAX_OPEN_CONNECTIONS" envDefault:"5"`
	AcquireTimeout       int  `env:"ACQUIRE_TIMEOUT" envDefault:"3"`
	PreferBinaryProtocol bool `env:"PREFER_BINARY_PROTOCOL" envDefault:"false"`
}

// Init config function that create config object end fill it from .env or env, if error then fatal
func InitConfig() *SConfig {
	var cfg SConfig
	absEnvPath := "$GOPATH/src/testServerStats/.env"
	if err := godotenv.Load(os.ExpandEnv(absEnvPath)); err != nil {
		fmt.Println("File .env not found, reading configuration from ENV")
	}

	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("Error init config, because: %v", err))
	}

	return &cfg
}
