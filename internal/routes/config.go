package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct {
	timeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"*"},
		MaxAge:           10,
	}).Handler(next)
}

func (c *Config) SetTimeout(timeInSeconds int) *Config {
	c.timeout = time.Duration(timeInSeconds) * time.Second
	return c
}

func (c *Config) GetTimeout() time.Duration {
	return c.timeout
}
