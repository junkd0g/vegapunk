package vegapunk

import "encoding/json"

const (
	enginesEndpoint = `/v1/engines`
)

type EnginesResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object      string      `json:"object"`
		ID          string      `json:"id"`
		Ready       bool        `json:"ready"`
		Owner       string      `json:"owner"`
		Created     interface{} `json:"created"`
		Permissions interface{} `json:"permissions"`
		Replicas    interface{} `json:"replicas"`
		MaxReplicas interface{} `json:"max_replicas"`
	} `json:"data"`
}

func (c client) GetEngines() (*EnginesResponse, error) {
	body, err := c.get(enginesEndpoint)
	if err != nil {
		return nil, err
	}

	var resp EnginesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
