package config

import (
	"time"
)

type Http struct {
	Host    string
	Agent   string
	Timeout time.Duration
}

func httpConfig(frequency time.Duration) Http {
	return Http{
		Host:    getEnv("HTTP_HOST", "localhost"),
		Agent:   getEnv("HTTP_AGENT", "THORmon Hlekkir"),
		Timeout: frequency * 9 / 10,
	}
}
