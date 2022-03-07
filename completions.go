package vegapunk

import (
	"encoding/json"
	"fmt"
)

const (
	completionsEndpoint = `/v1/engines/%s/completions`
)

type CompletionsRequest struct {
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
	BestOf           int     `json:"best_of"`
	Echo             bool    `json:"echo"`
	Logprobs         int     `json:"logprobs"`
	Stream           bool    `json:"stream"`
}

func (c client) CreateCompletions(engineID string, cr CompletionsRequest) (interface{}, error) {
	endpoint := fmt.Sprintf(completionsEndpoint, engineID)
	request, err := json.Marshal(cr)

	body, err := c.post(endpoint, string(request))
	if err != nil {
		return nil, err
	}

	var resp interface{}
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
