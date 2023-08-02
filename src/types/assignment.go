package types;

type AssignmentGroup struct {
    AmbientID string        `json:"ambient_id"`;
    Professor Professor     `json:"professor"`;
    Enrolled string         `json:"enrolled"`;

    Schedules []struct{
        Day int             `json:"day"`;
        ClassType string    `json:"type"`;
        From string         `json:"from"`;
        To string           `json:"to"`;
    }                       `json:"schedules"`;
}

type Assignment struct {
    Code string                 `json:"code"`;
    Name string                 `json:"name"`;
    Credits string              `json:"credits"`;
    Groups []AssignmentGroup    `json:"groups"`;
}

