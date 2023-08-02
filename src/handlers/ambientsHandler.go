package handlers

import (
	"ambients-end/src/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
    var assignments []types.Assignment;
    var ambientsList []types.Ambient;

    assignmentsFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/assignments.json");
    if (openErr != nil) { log.Fatal(openErr); }
    decodeErr := json.NewDecoder(assignmentsFile).Decode(&assignments);
    if (decodeErr != nil) { log.Fatal(decodeErr); }
    var sampleAmbient types.Ambient;
    for _, assignment := range assignments {
        for _, e := range assignment.Groups {
            sampleAmbient.AmbientID = e.AmbientID;
            sampleAmbient.Category = strings.Split(e.AmbientID, ".")[0];
            sampleAmbient.Description = "";
            sampleAmbient.Gallery = nil;
            sampleAmbient.Tags = nil;

            ambientsList = append(ambientsList, sampleAmbient);
        }
    }

    urlPath := filepath.Clean((*r).URL.Path);
    log.Println((*r).Method + " @ " + urlPath);

    pathElements := strings.Split(urlPath, "/");

    log.Println("-> " + fmt.Sprint(len(pathElements)));

    if (len(pathElements) < 4) {
        DisplayAmbients(w, ambientsList);
    } else if (len(pathElements) == 5 && pathElements[3] == "category") {
        displayCategory := getCategoryList(ambientsList, pathElements[4]);
        DisplayAmbients(w, displayCategory);
    } else if (pathElements[3] == "ambient_id") {
        log.Println(r.URL.RawQuery)
        ambientID := r.URL.Query().Get("id");
        displayList := getIdList(ambientsList, ambientID);
        DisplayAmbients(w, displayList);
    } else {
        DisplayAmbients(w, ambientsList);
    }
}

func DisplayAmbients(w http.ResponseWriter, ambientsList []types.Ambient) {
    ambientsStringData := new(strings.Builder);
    json.NewEncoder(ambientsStringData).Encode(ambientsList);
    w.Write([]byte(ambientsStringData.String()));
}

func getCategoryList(ambientsList []types.Ambient, category string) []types.Ambient {
    var newAmbientsList []types.Ambient;
    for _, ambient := range ambientsList {
        if (ambient.Category == category) {
            newAmbientsList = append(newAmbientsList, ambient);
        }
    }
    return newAmbientsList;
}

func getIdList(ambientsList []types.Ambient, id string) []types.Ambient {
    var newAmbientsList []types.Ambient;
    for _, ambient := range ambientsList {
        if (ambient.AmbientID == id) {
            newAmbientsList = append(newAmbientsList, ambient);
        }
    }
    return newAmbientsList;
}

