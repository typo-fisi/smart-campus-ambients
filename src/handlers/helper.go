package handlers

import (
	openaiClient "ambients-end/src/openai/client"
	openaiTypes "ambients-end/src/openai/types"
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

    var incomingBody types.IncomingMsg;
    incomingErr := json.NewDecoder((*r).Body).Decode(&incomingBody);

    if (incomingErr != nil) {
        w.WriteHeader(500);
        log.Println(incomingErr);
        return;
    }

    inferredCategory, categoryErr := inferCategory(incomingBody);
    if (categoryErr != nil) {
        log.Println(categoryErr);
        w.WriteHeader(500);
        return;
    }

    inferredId, idErr := inferId(incomingBody, strings.ToLower(inferredCategory));
    if (idErr != nil) {
        log.Println(idErr);
        w.WriteHeader(500);
        return;
    }

    serverResp := struct{
        Category string `json:"category"`;
        ID string `json:"id"`;
    }{
        Category: inferredCategory,
        ID: inferredId,
    }
    respStr := new(strings.Builder);

    finalDecodeErr := json.NewEncoder(respStr).Encode(serverResp);
    if (finalDecodeErr != nil) {
        log.Println(finalDecodeErr);
        w.WriteHeader(500);
        return;
    }

    w.Write([]byte(respStr.String()));
}

func inferCategory(income types.IncomingMsg) (string, error) {
    var promptStr string;
    var response types.IncomingResp;

    sample1, sample1Err := os.ReadFile("./dump/sample1");
    if (sample1Err != nil){ log.Fatal(sample1Err); }

    promptStr += string(sample1) + income.Context;

    message := openaiTypes.Message{
        Role: "system",
        Content: promptStr,
    }

    getCategoryRequest := openaiTypes.Request{
        Model: "gpt-3.5-turbo",
        Messages: []openaiTypes.Message{
            message,
            {
                Role: "user",
                Content: income.Prompt,
            },
        },
        Temperature: 0.7,
    }

    getCategoryResponse, _ := openaiClient.GetResponse(getCategoryRequest);
    responseDecodeErr := json.NewDecoder(getCategoryResponse.Body).Decode(&response);

    if (responseDecodeErr != nil) {
        log.Println(responseDecodeErr);
        return "", responseDecodeErr;
    }

    return response.Choices[0].Message.Content, nil;
}

func inferId(income types.IncomingMsg, category string) (string, error) {
    var response types.IncomingResp;

    listFile, openErr := os.Open("./dump/ambients.json");
    if (openErr != nil) {
        log.Println(openErr);
        return "", openErr;
    }

    var ambients []types.Ambient;

    decodeErr := json.NewDecoder(listFile).Decode(&ambients);
    if (decodeErr != nil) {
        log.Println(decodeErr);
        return "", decodeErr;
    }

    var list []prompt.ListElem;
    var sample prompt.ListElem;

    for _, ambient := range ambients {
        if (ambient.Category != category) { continue; }

        sample.Name = ambient.Name;
        sample.Description = ambient.Description;
        sample.ID = ambient.AmbientID;

        list = append(list, sample);
    }

    respStr := new(strings.Builder);

    respGenErr := json.NewEncoder(respStr).Encode(list);
    if (respGenErr != nil) {
        log.Println(respGenErr);
        return "", respGenErr;
    }

    sample2, sample2Err := os.ReadFile("./dump/sample2");
    if (sample2Err != nil){
        log.Println(sample2Err);
        return "", sample2Err;
    }

    promptStr := string(sample2) + respStr.String() + "\n" + income.Context;
    log.Println(promptStr);

    message := openaiTypes.Message{
        Role: "system",
        Content: promptStr,
    }

    getCategoryRequest := openaiTypes.Request{
        Model: "gpt-3.5-turbo",
        Messages: []openaiTypes.Message{
            message,
            {
                Role: "user",
                Content: income.Prompt,
            },
        },
        Temperature: 0.3,
    }

    getIdResponse, _ := openaiClient.GetResponse(getCategoryRequest);
    responseDecodeErr := json.NewDecoder(getIdResponse.Body).Decode(&response);

    if (responseDecodeErr != nil) {
        log.Println(responseDecodeErr);
        return "", responseDecodeErr;
    }

    return response.Choices[0].Message.Content, nil;
}

