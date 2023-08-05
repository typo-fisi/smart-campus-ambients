package types;

type Message struct {
    Role string     `json:"role"`;
    Content string  `json:"content"`;
}

type Request struct {
    Model string        `json:"model"`;
    Messages []Message  `json:"messages"`;
    Temperature float32 `json:"temperature"`;
}

