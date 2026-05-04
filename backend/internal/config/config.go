package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port               string   `envconfig:"PORT" default:"8080"`
	DatabaseUser       string   `envconfig:"DATABASE_USER" required:"true"`
	DatabasePassword   string   `envconfig:"DATABASE_PASSWORD" required:"true"`
	DatabaseName       string   `envconfig:"DATABASE_NAME" required:"true"`
	DatabaseHost       string   `envconfig:"DATABASE_HOST" required:"true"`
	DatabasePort       string   `envconfig:"DATABASE_PORT" default:"5432"`
	JwtSecret          string   `envconfig:"jWT"`
	CorsAllowedOrigins []string `envconfig:"CORS_ALLOWED_ORIGINS" default:"http://localhost:5173,http://localhost:3000"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName)
}
