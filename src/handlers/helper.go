package handlers

import (
	openaiClient "ambients-end/src/openai/client"
	openaiTypes "ambients-end/src/openai/types"
	"ambients-end/src/prompt"
	"ambients-end/src/types"
	"encoding/json"
	"fmt"
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

    var incomingBody types.IncomingMsg;
    incomingErr := json.NewDecoder((*r).Body).Decode(&incomingBody);

    if (incomingErr != nil) {
        w.WriteHeader(500);
        return;
    }
}

func inferCategory(income types.IncomingMsg) (string, error) {
    var promptStr string;

    sample1, sample1Err := os.ReadFile("./dump/sample1");
    if (sample1Err != nil){ log.Fatal(sample1Err); }

    promptStr += string(sample1) + income.Context;

    message := openaiTypes.Message{
        Role: "system",
        Content: promptStr,
    }

    textazo1 := openaiTypes.Request{
        Model: "gpt-3.5-turbo",
        Messages: []openaiTypes.Message{
            message,
            {
                Role: "user",
                Content: income.Prompt,
            },
        },
    }

    resp1, _ := openaiClient.GetResponse(textazo1);
    fmt.Println(resp1);
    return "", nil;
}

func inferId(income types.IncomingMsg) {
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

    respStr := new(strings.Builder);

    respGenErr := json.NewEncoder(respStr).Encode(list);
    if (respGenErr != nil) { log.Fatal(respGenErr); }
}

