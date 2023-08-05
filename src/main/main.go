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
    dumpfilename := "./dump/assignments.json"

    holaFile, openErr := os.Open(dumpfilename);
    if (openErr != nil) { log.Fatal(openErr); }

    var holaSi []types.Assignment
    json.NewDecoder(holaFile).Decode(&holaSi);

    for i, e := range holaSi {
        for j := range e.Groups {
            chooser := rand.Intn(12) + 1;

            var id int;
            var min int;
            var max int;

            switch (chooser) {
            case 1:
                min = 315;
                max = 327;
                id = rand.Intn(max - min) + min;
            case 2:
                min = 217;
                max = 220;
                id = rand.Intn(max - min) + min;
            case 3:
                min = 223;
                max = 224;
                id = rand.Intn(max - min) + min;
            case 4:
                min = 228;
                max = 228;
                id = 228;
            case 5:
                min = 201;
                max = 202;
                id = rand.Intn(max - min) + min;
            case 6:
                min = 211;
                max = 212;
                id = rand.Intn(max - min) + min;
            case 7:
                min = 230;
                max = 231;
                id = rand.Intn(max - min) + min;
            case 8:
                min = 208;
                max = 209;
                id = rand.Intn(max - min) + min;
            case 9:
                min = 328;
                max = 328;
                id = 328;
            case 10:
                min = 301;
                max = 301;
                id = 301;
            case 11:
                min = 109;
                max = 118;
                id = rand.Intn(max - min) + min;
            case 12:
                min = 123;
                max = 125;
                id = rand.Intn(max - min) + min;
            case 13:
                min = 127;
                max = 135;
                id = rand.Intn(max - min) + min;
            default:
                min = 0;
                max = 0;
                id = 0;
            }

            holaSi[i].Groups[j].AmbientID = id;
        }
    }
    holastr := new(strings.Builder)
    json.NewEncoder(holastr).Encode(holaSi)

    fmt.Println(holastr);
}
