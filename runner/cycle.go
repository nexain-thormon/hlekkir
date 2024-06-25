package runner

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"hlekkir/config"
	"hlekkir/fetcher"
	"hlekkir/logger"
	"hlekkir/olgerd"

	"github.com/redis/go-redis/v9"
)

func cycle(ctx context.Context, cfg config.App, rdb *redis.Client, httpClient *http.Client) {
	if err := cleanup(ctx, cfg, rdb); err != nil {
		logger.Log.Error().Err(err).Msg("Error cleaning up expired data")
		return
	}

	nodes, err := olgerd.Nodes(ctx, rdb, cfg.Redis.Olgerd)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error retrieving node data")
		return
	}

	resultsCh := make(chan map[string]interface{}, len(nodes))
	doneCh := make(chan struct{})

	var wg sync.WaitGroup

	for _, n := range nodes {
		wg.Add(1)
		go func(n olgerd.Node) {
			defer wg.Done()
			fetcher.Status(ctx, httpClient, n, cfg, resultsCh)
		}(n)
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	results := collection(resultsCh, doneCh)

	combinedResults, err := json.Marshal(results)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error marshalling combined results")
		return
	}

	if err := rdb.Set(ctx, cfg.Redis.Hlekkir, combinedResults, 0).Err(); err != nil {
		logger.Log.Error().Err(err).Msg("Error setting combined results in Redis")
	}
}
