package types;

type AmbientInterface interface {
    GetAmbientId() string;
}

type Ambient struct{
    Name string         `json:"name"`;
    Category string     `json:"category"`;
    AmbientID int       `json:"ambient_id"`;
    Description string  `json:"description"`;
    Gallery []struct{
        Src string      `json:"src"`;
    }                   `json:"gallery"`;
    Tags []string       `json:"tags"`;
}

func (sa Ambient) GetAmbientId() int { return sa.AmbientID; }

type SalonAmbient struct {
    Ambient;
    Assignments []AssignmentGroup;
}

func (sa SalonAmbient) GetAmbientId() int { return sa.AmbientID; }


