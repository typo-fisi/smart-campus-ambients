package types

import "ambients-end/src/openai/types"

type IncomingResp struct {
    Id string                   `json:"int"`;
    Object string               `json:"object"`;
    Created interface{}         `json:"created"`;
    Model string                `json:"model"`;
    Choices []struct{
        Index int               `json:"index"`;
        Message types.Message   `json:"message"`;
        FinishReason string     `json:"finish_reason"`;
    }                           `json:"choices"`;
    Usage struct{
        PromptTokens int        `json:"prompt_tokens"`;
        CompletionTokens int    `json:"completion_tokens"`;
        TotalTokens int         `json:"total_tokens"`;
    };
}

