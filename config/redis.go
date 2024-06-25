package config

type Redis struct {
	Url     string
	Hlekkir string
	Olgerd  string
	Cleanup int64
}

func redisConfig(isProd bool) Redis {
	url := getEnv("REDIS_URL", "")
	if url == "" {
		url = "redis://localhost:6379/0"
		if isProd {
			url = "redis://redis:6379/0"
		}
	}

	return Redis{
		Url:     url,
		Hlekkir: getEnv("REDIS_HLEKKIR", "thormon_hlekkir_mainnet"),
		Olgerd:  getEnv("REDIS_OLGERD", "thormon_olgerd_mainnet_hlekkir"),
		Cleanup: getEnvAsInt64("REDIS_CLEANUP", 300),
	}
}
