package olgerd

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Node struct {
	Address string `json:"node_address"`
	IP      string `json:"ip_address"`
}

func Nodes(ctx context.Context, rdb *redis.Client, key string) ([]Node, error) {
	jsonData, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var nodes []Node
	if err := json.Unmarshal([]byte(jsonData), &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}
