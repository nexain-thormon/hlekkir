package runner

import (
	"context"
	"encoding/json"
	"time"

	"hlekkir/config"

	"github.com/redis/go-redis/v9"
)

func cleanup(ctx context.Context, cfg config.App, rdb *redis.Client) error {
	data, err := rdb.Get(ctx, cfg.Redis.Hlekkir).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	hlekkirData, err := unmarshal(data)
	if err != nil {
		return err
	}

	flushExpired(hlekkirData, cfg.Redis.Cleanup)

	updatedData, err := json.Marshal(hlekkirData)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, cfg.Redis.Hlekkir, updatedData, 0).Err()
}

func unmarshal(data string) (map[string]map[string]interface{}, error) {
	var hlekkirData map[string]map[string]interface{}
	return hlekkirData, json.Unmarshal([]byte(data), &hlekkirData)
}

func flushExpired(hlekkirData map[string]map[string]interface{}, cleanupThreshold int64) {
	currentTime := time.Now().Unix()
	for key, value := range hlekkirData {
		if timestamp, ok := value["timestamp"].(int64); ok && currentTime-timestamp > cleanupThreshold {
			delete(hlekkirData, key)
		}
	}
}
