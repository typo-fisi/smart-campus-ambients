package handlers

import (
	"ambients-end/src/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func AmbientsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json; charset=utf-8");
    EnableCors(&w);

    if ((*r).Method == "OPTIONS") {
        w.WriteHeader(200);
        w.Write([]byte("{}"))
    }

    var assignments []types.Assignment;
    var ambientsList []types.Ambient;

    assignmentsFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/assignments2.json");
    if (openErr != nil) { log.Fatal(openErr); }

    decode_assignmentsErr := json.NewDecoder(assignmentsFile).Decode(&assignments);
    if (decode_assignmentsErr != nil) { log.Fatal(decode_assignmentsErr); }

    ambientsFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/ambients.json");
    if (openErr != nil) { log.Fatal(openErr); }

    decode_ambientsErr := json.NewDecoder(ambientsFile).Decode(&ambientsList);
    if (decode_ambientsErr != nil) { log.Fatal(decode_ambientsErr); }


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
        ambientBigInt, parseErr := strconv.ParseInt(r.URL.Query().Get("id"), 10, 0);
        ambientID := int(ambientBigInt);

        if (parseErr != nil) {
            displayAmbients(w, nil);
            log.Println(parseErr);
        }

        switch getCategoryById(ambientsList, ambientID) {
            case "salones":
                displayList := getIdListSalones(ambientsList, ambientID, assignments);
                displayAmbientsSalones(w, displayList);
            default:
                displayList := getIdList(ambientsList, ambientID, assignments);
                displayAmbients(w, displayList);
        }

    } else {
        displayAmbients(w, ambientsList);
    }
}

func displayAmbients(w http.ResponseWriter, ambientsList []types.Ambient) {
    ambientsStringData := new(strings.Builder);
    json.NewEncoder(ambientsStringData).Encode(ambientsList);
    w.Write([]byte(ambientsStringData.String()));
}

func displayAmbientsSalones(w http.ResponseWriter, ambientsList []types.SalonAmbient) {
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

func getIdList(ambientsList []types.Ambient, id int, assignments []types.Assignment) []types.Ambient {
    var newAmbientsList []types.Ambient;
    for _, ambient := range ambientsList {
        if (ambient.AmbientID == id) {
            newAmbientsList = append(newAmbientsList, ambient);
        }
    }
    return newAmbientsList;
}

func getCategoryById(ambientsList []types.Ambient, id int) string {
    for _, ambient := range ambientsList {
        if (ambient.AmbientID == id) {
            return ambient.Category;
        }
    }
    return "";
}

func getIdListSalones(ambientsList []types.Ambient, id int, assignments []types.Assignment) []types.SalonAmbient {
    var classroomsList []types.SalonAmbient;
    var sampleClassroom types.SalonAmbient;

    for _, ambient := range ambientsList {
        if (ambient.AmbientID == id) {
            if (ambient.Category == "salones") {
                sampleClassroom.Ambient = ambient;
                sampleClassroom.Assignments = getAssociatedAssignmentGroups(id, assignments);
                classroomsList = append(classroomsList, sampleClassroom);
            }
        }
    }
    return classroomsList;
}


func getAssociatedAssignmentGroups(id int, assignments []types.Assignment) []types.AssignmentGroup {
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



