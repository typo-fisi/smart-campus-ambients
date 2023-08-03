package main

import (
	"ambients-end/src/handlers"
	"ambients-end/src/types"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func main2() {
    holaFile, openErr := os.Open("/home/bauer/Projects/smartcampus-ambients-end/dump/assignments.json");
    if (openErr != nil) { log.Fatal(openErr); }

    var holaSi []types.Assignment
    json.NewDecoder(holaFile).Decode(&holaSi);

    for i, e := range holaSi {
        for j := range e.Groups {
            holaSi[i].Groups[j].AmbientID = rand.Intn(34) + 1;
        }
    }
    holastr := new(strings.Builder)
    json.NewEncoder(holastr).Encode(holaSi)

    fmt.Println(holastr);

}




func main() {
    http.HandleFunc("/api/assignments/", handlers.AssignmentsHandler);
    http.HandleFunc("/api/ambients/", handlers.AmbientsHandler);
    http.HandleFunc("/api/professors/", handlers.ProfessorHandler);
    http.ListenAndServe(":4000", nil);
}


