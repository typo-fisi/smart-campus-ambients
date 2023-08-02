package types;

type AssignmentGroup struct {
    AmbientID string        `json:"ambient_id"`;
    Group string            `json:"group"`;
    Professor Professor     `json:"professor"`;
    Enrolled string         `json:"enrolled"`;

    Schedules []struct{
        Day string          `json:"day"`;
        ClassType string    `json:"classtype"`;
        From int            `json:"from"`;
        To int              `json:"to"`;
    }                       `json:"schedule"`;
}

type Assignment struct {
    Code string                 `json:"code"`;
    Name string                 `json:"name"`;
    Credits string              `json:"credits"`;
    Groups []AssignmentGroup    `json:"groups"`;
}

