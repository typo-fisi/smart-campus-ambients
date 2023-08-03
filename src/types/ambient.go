package types

type AmbientInterface interface {
    GetAmbientId() string;
}

type Ambient struct{
    Category string     `json:"category"`;
    AmbientID string    `json:"ambient_id"`;
    Description string  `json:"description"`;
    Gallery []struct{
        Src string      `json:"src"`;
    }                   `json:"gallery"`;
    Tags []string       `json:"tags"`;
}

type SalonAmbient struct {
    Ambient;
    Assignments []AssignmentGroup;
}

func (sa SalonAmbient) GetAmbientId() string {
    return sa.AmbientID;
}


