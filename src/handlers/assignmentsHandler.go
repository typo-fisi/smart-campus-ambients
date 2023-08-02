package handlers

import (
	"ambients-end/src/types"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func separatePath(path string) []string {
    return strings.Split(path, "/");
}

func AssignmentsHandler(w http.ResponseWriter, r *http.Request) {
    assignmentsFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/assignments.json");
    if (openErr != nil) { log.Fatal(openErr) }

    urlPath := (*r).URL.Path;
    pathMembers := separatePath(urlPath);

    var allAssignments []types.Assignment;

    decoder := json.NewDecoder(assignmentsFile);
    decodingErr := decoder.Decode(&allAssignments);
    if (decodingErr != nil) { log.Fatal(decodingErr) }

    var pathID string;

    if (pathMembers[3] != "") {
        pathID = pathMembers[3];
        getSingleAssignment(w, allAssignments, pathID);
    } else {
        getAllAssignments(w, allAssignments);
    }
}

func getAllAssignments(w http.ResponseWriter, assignments []types.Assignment) {
    response := new(strings.Builder);

    encoder := json.NewEncoder(response);
    err := encoder.Encode(assignments);
    if (err != nil) { log.Fatal(err); }

    w.Write([]byte(response.String()));
}

func getSingleAssignment(w http.ResponseWriter, assignments []types.Assignment, ID string) {
    response := new(strings.Builder);
    encoder := json.NewEncoder(response);

    for _, assignment := range assignments {
        if (assignment.Code == ID) {
            encoder.Encode(assignment);
            w.Write([]byte(response.String()));
        }
    }
}






