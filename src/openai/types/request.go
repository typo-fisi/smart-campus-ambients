package types;

type Request struct {
    Model string    `json:"model"`;
    Prompt string   `json:"prompt"`;
    MaxTokens int   `json:"max_tokens"`;
    Temperature int `json:"temperature"`;
}








