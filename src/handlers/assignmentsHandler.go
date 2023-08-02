package handlers

import (
	"ambients-end/src/types"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func AssignmentsHandler(w http.ResponseWriter, r *http.Request) {
    var allAssignments []types.Assignment;
    assignmentsFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/assignments.json");

    if (openErr != nil) { log.Fatal(openErr) }
    decoder := json.NewDecoder(assignmentsFile);

    decodingErr := decoder.Decode(&allAssignments);
    if (decodingErr != nil) { log.Fatal(decodingErr) }

    (w).Write([]byte(allAssignments[0].Code));
}
