package main

import (
	"context"
	"net/http"

	"hlekkir/client"
	"hlekkir/config"
	"hlekkir/runner"
)

func main() {
	cfg := config.Load()
	rdb := client.Redis(cfg.Redis.Url)
	httpClient := &http.Client{
		Timeout: cfg.Http.Timeout,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runner.Start(ctx, cfg, rdb, httpClient)
}
