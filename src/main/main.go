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

func main() {
    http.HandleFunc("/", handlers.GeneralHandle);
    http.HandleFunc("/api/assignments/", handlers.AssignmentsHandler);
    http.HandleFunc("/api/ambients/", handlers.AmbientsHandler);
    http.HandleFunc("/api/helper", handlers.Helper);
    http.ListenAndServe(":4000", nil);
}

func main2() {
    var dumpfilename string;

    //dumpfilenamme = {name}

    holaFile, openErr := os.Open("./dump/" + dumpfilename);
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
