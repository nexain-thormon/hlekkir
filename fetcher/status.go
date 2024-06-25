package fetcher

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"hlekkir/config"
	"hlekkir/logger"
	"hlekkir/olgerd"
)

func Status(ctx context.Context, httpClient *http.Client, node olgerd.Node, cfg config.App, resultsCh chan<- map[string]interface{}) {
	url := "http://" + node.IP + ":6040/status/scanner"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logger.Log.Error().Err(err).Str("ip", node.IP).Msg(node.Address)
		return
	}
	req.Header.Set("Host", cfg.Http.Host)
	req.Header.Set("User-Agent", cfg.Http.Agent)

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Log.Error().Err(err).Str("ip", node.IP).Msg(node.Address)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Log.Error().Str("ip", node.IP).Int("status", resp.StatusCode).Msg(node.Address)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error().Err(err).Msg(node.Address)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		logger.Log.Error().Err(err).Msg(node.Address)
		return
	}

	result["timestamp"] = time.Now().Unix()

	logger.Log.Info().Str("ip", node.IP).Interface("probe", result).Msg(node.Address)
	resultsCh <- map[string]interface{}{node.Address: result}
}
