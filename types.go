package vegapunk

import "time"

type CompletionsResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Model   string    `json:"model"`
	Created time.Time `json:"created"`
	Choices []Choice  `json:"choices"`
}

type Choice struct {
	Text         string   `json:"text"`
	FinishReason string   `json:"finish_reason"`
	Index        int      `json:"index"`
	Logprobs     Logprobs `json:"logprobs"`
}

type Logprobs struct {
	Tokens        []string   `json:"tokens"`
	CompleteToken string     `json:"complete"`
	Choices       []LPChoice `json:"choices"`
}

type LPChoice struct {
	Text    string  `json:"text"`
	Logprob float64 `json:"logprob"`
}
