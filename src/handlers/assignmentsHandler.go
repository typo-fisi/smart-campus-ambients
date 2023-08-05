package handlers

import (
	"ambients-end/src/types"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func separatePath(path string) []string {
    return strings.Split(path, "/");
}

func AssignmentsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json; charset=utf-8");
    EnableCors(&w);

    assignmentsFile, openErr := os.Open("./dump/assignments2.json");
    if (openErr != nil) { log.Fatal(openErr) }

    urlPath := filepath.Clean((*r).URL.Path);
    pathMembers := separatePath(urlPath);

    var allAssignments []types.Assignment;

    decoder := json.NewDecoder(assignmentsFile);
    decodingErr := decoder.Decode(&allAssignments);
    if (decodingErr != nil) { log.Fatal(decodingErr) }

    var pathID string;

    if (len(pathMembers) < 4) {
        writeAssignments(w, allAssignments);
    } else {
        pathID = pathMembers[3];
        displayAssignments := getAssignmentsById(pathID, allAssignments);
        writeAssignments(w, displayAssignments);
    }
}

func writeAssignments(w http.ResponseWriter, assignments []types.Assignment) {
    response := new(strings.Builder);

    encoder := json.NewEncoder(response);
    err := encoder.Encode(assignments);
    if (err != nil) { log.Fatal(err); }

    w.Write([]byte(response.String()));
}

func getAssignmentsById(ID string, assignments []types.Assignment) []types.Assignment {
    var displayAssignments []types.Assignment;

    for _, assignment := range assignments {
        if (assignment.Code == ID) {
            displayAssignments = append(displayAssignments, assignment);
        }
    }

    return displayAssignments;
}






