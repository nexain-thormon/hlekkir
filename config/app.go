package config

import (
	"time"
)

type App struct {
	Environment string
	IsProd      bool
	AppCycle    int
	Http        Http
	Redis       Redis
	Frequency   time.Duration
}

func Load() App {
	environment := getEnv("APP_ENV", "development")
	isProd := environment == "production"
	frequency := time.Duration(getEnvAsInt("FREQUENCY", 20)) * time.Second

	return App{
		Environment: environment,
		IsProd:      isProd,
		AppCycle:    getEnvAsInt("APP_CYCLE", -1),
		Http:        httpConfig(frequency),
		Redis:       redisConfig(isProd),
		Frequency:   frequency,
	}
}
