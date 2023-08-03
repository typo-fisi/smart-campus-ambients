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
        displayAmbients(w, ambientsList);
    } else if (len(pathElements) == 5 && pathElements[3] == "category") {
        displayCategory := getCategoryList(ambientsList, pathElements[4]);
        displayAmbients(w, displayCategory);
    } else if (pathElements[3] == "ambient_id") {
        log.Println(r.URL.RawQuery)
        ambientID := r.URL.Query().Get("id");
        displayList := getIdList(ambientsList, ambientID, assignments);
        displayAmbientsInterface(w, displayList);
    } else {
        displayAmbients(w, ambientsList);
    }
}

func displayAmbients(w http.ResponseWriter, ambientsList []types.Ambient) {
    ambientsStringData := new(strings.Builder);
    json.NewEncoder(ambientsStringData).Encode(ambientsList);
    w.Write([]byte(ambientsStringData.String()));
}

func displayAmbientsInterface(w http.ResponseWriter, ambientsList []types.AmbientInterface) {
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

func getIdList(ambientsList []types.Ambient, id string, assignments []types.Assignment) []types.AmbientInterface {
    var newAmbientsList []types.Ambient;
    var classroomsList []types.AmbientInterface;

    var sampleClassroom types.SalonAmbient;

    for _, ambient := range ambientsList {
        if (ambient.AmbientID == id) {
            newAmbientsList = append(newAmbientsList, ambient);

            if (ambient.Category == "salones") {
                sampleClassroom.Ambient = ambient;
                sampleClassroom.Assignments = getAssociatedAssignmentGroups(id, assignments);
                classroomsList = append(classroomsList, sampleClassroom);
            }
        }
    }

    return classroomsList;
}

func getAssociatedAssignmentGroups(id string, assignments []types.Assignment) []types.AssignmentGroup {
    var displayGroups []types.AssignmentGroup;

    for _, assignment := range assignments {
        for _, group := range assignment.Groups {
            if (group.AmbientID == id) {
                displayGroups = append(displayGroups, group);
            }
        }
    }

    return displayGroups;
}




