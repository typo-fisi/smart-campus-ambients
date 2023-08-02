package handlers

import (
	"ambients-end/src/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
type Ambient struct{
    Category string
    ID string;
    Location struct{
        piso int;
        pavilion string;
    };
    Description string;
    Fotos []string;
    Tags []string;
}

type AssignmentGroup struct {
    AmbientID string        `json:"ambient_id"`;
    Group string            `json:"group"`;
    Professor Professor     `json:"professor"`;
    Enrolled string         `json:"enrolled"`;

    Schedules []struct{
        Day string          `json:"day"`;
        ClassType string    `json:"type"`;
        From int            `json:"from"`;
        To int              `json:"to"`;
    }                       `json:"schedule"`;
}

type Assignment struct {
    Code string                 `json:"code"`;
    Name string                 `json:"name"`;
    Credits string              `json:"credits"`;
    Groups []AssignmentGroup    `json:"groups"`;
}*/

func AmbientsHandler(w http.ResponseWriter, r *http.Request) {
    //var ambientsList []types.Ambient;

    //in the mean time, we will get the ambients' data
    //from the assignments' data and some random stuff
    var assignments []types.Assignment;
    assignmentsFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/assignments.json");
    if (openErr != nil) { log.Fatal(openErr); }
    decodeErr := json.NewDecoder(assignmentsFile).Decode(&assignments);
    if (decodeErr != nil) { log.Fatal(decodeErr); }


    fmt.Println(log.Ltime);
}

