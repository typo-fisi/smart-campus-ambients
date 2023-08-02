package types

type Ambient struct{
    Category string     `json:"category"`;
    AmbientID string   `json:"ambient_id"`;
    Description string  `json:"description"`;
    Gallery []struct{
        Src string      `json:"src"`;
    }                   `json:"gallery"`;
    Tags []string       `json:"tags"`;
}

