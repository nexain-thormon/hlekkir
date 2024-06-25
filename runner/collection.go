package runner

func collection(resultsCh <-chan map[string]interface{}, doneCh <-chan struct{}) map[string]interface{} {
	results := make(map[string]interface{})
	for {
		select {
		case result := <-resultsCh:
			for k, v := range result {
				results[k] = v
			}
		case <-doneCh:
			return results
		}
	}
}
