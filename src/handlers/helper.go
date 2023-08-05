package handlers

import (
	"ambients-end/src/prompt"
	"ambients-end/src/types"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func Helper(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json; charset=utf-8");
    EnableCors(&w);
    if ((*r).Method == "OPTIONS") {
        EnableCors(&w);
        w.WriteHeader(200);
    }


    listFile, openErr := os.Open("./dump/ambients.json");
    if (openErr != nil) { log.Fatal(openErr); }

    var ambients []types.Ambient;

    decodeErr := json.NewDecoder(listFile).Decode(&ambients);
    if (decodeErr != nil) { log.Fatal(openErr); }

    var list []prompt.ListElem;
    var sample prompt.ListElem;

    for _, ambient := range ambients {
        sample.Name = ambient.Name;
        sample.Description = ambient.Description;
        sample.ID = ambient.AmbientID;

        list = append(list, sample);
    }

    resp := new(strings.Builder);

    respGenErr := json.NewEncoder(resp).Encode(list);
    if (respGenErr != nil) { log.Fatal(respGenErr); }

    w.Write([]byte(resp.String()));
}
