package vegapunk

import (
	"encoding/json"
	"fmt"
)

const (
	enginesEndpoint = `/v1/engines`
)

/*
	Response example for the get endpoint /v1/engines

	{
        "object": "list",
        "data": [
            {
               "object": "engine",
               "id": "ada-code-search-text",
                "ready": true,
                "owner": "openai",
                "created": null,
                "permissions": null,
                "replicas": null,
                "max_replicas": null
            },
            {
                "object": "engine",
                "id": "text-similarity-davinci-001",
                "ready": true,
                "owner": "openai",
                "created": null,
                "permissions": null,
                "replicas": null,
                "max_replicas": null
            }
        ]
    }
*/

// EnginesResponse reponse for /v1/engines
type EnginesResponse struct {
	Object string           `json:"object"`
	Data   []EngineResponse `json:"data"`
}

// EngineResponse reponse for /v1/engine/"id"
type EngineResponse struct {
	Object      string      `json:"object"`
	ID          string      `json:"id"`
	Ready       bool        `json:"ready"`
	Owner       string      `json:"owner"`
	Created     interface{} `json:"created"`
	Permissions interface{} `json:"permissions"`
	Replicas    interface{} `json:"replicas"`
	MaxReplicas interface{} `json:"max_replicas"`
}

// GetEngines return an array of the openai engines
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

// GetEngine return specific openai engine
func (c client) GetEngine(engineID string) (*EngineResponse, error) {
	engineEndpoint := fmt.Sprintf("%s/%s", enginesEndpoint, engineID)
	body, err := c.get(engineEndpoint)
	if err != nil {
		return nil, err
	}

	var resp EngineResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
