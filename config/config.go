package config

import "github.com/caarlos0/env/v6"

type Config struct {
	HTTPServer HTTPServerConfig
	Logger     LoggerConfig
	Mongo      MongoConfig
	JWT        JWTConfig
	Encrypter  EncrypterConfig
}

type MongoConfig struct {
	URI    string `env:"MONGO_URI" envDefault:"mongodb://mongo:password@localhost:27117"`
	DBName string `env:"MONGO_DB_NAME" envDefault:"tanca-event-mongo"`
}

type HTTPServerConfig struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Mode string `env:"MODE" envDefault:"development"`
}

type LoggerConfig struct {
	Level    string `env:"LOGGER_LEVEL" envDefault:"debug"`
	Mode     string `env:"MODE" envDefault:"development"`
	Encoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET"`
}

type EncrypterConfig struct {
	Key string `env:"ENCRYPT_KEY"`
}

// Load loads the configuration from the environment variables.
func Load() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
