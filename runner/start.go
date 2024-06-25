package runner

import (
	"context"
	"net/http"
	"time"

	"hlekkir/config"

	"github.com/redis/go-redis/v9"
)

func Start(ctx context.Context, cfg config.App, rdb *redis.Client, httpClient *http.Client) {
	ticker := time.NewTicker(cfg.Frequency)
	defer ticker.Stop()
	counter := 0

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			counter++
			if !cfg.IsProd && cfg.AppCycle != -1 && counter > cfg.AppCycle {
				return
			}
			cycle(ctx, cfg, rdb, httpClient)
		}
	}
}
