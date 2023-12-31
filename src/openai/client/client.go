package openai

import (
	"ambients-end/src/openai/types"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func newOpenAIReq(body types.Request) *http.Request {
    openAiPath := "https://api.openai.com/v1/chat/completions";
    openAiKey := "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";

    bodyStr := new(strings.Builder);
    encodeErr := json.NewEncoder(bodyStr).Encode(body);
    if (encodeErr != nil) { log.Fatal(encodeErr); }

    log.Println(bodyStr);

    openAiReq, reqErr := http.NewRequest(http.MethodPost, openAiPath, bytes.NewReader([]byte(bodyStr.String())));
    if (reqErr != nil) { log.Fatal(reqErr); }

    (*openAiReq).Header.Set("Content-Type", "application/json");
    (*openAiReq).Header.Set("Authorization", "Bearer " + openAiKey);

    return openAiReq;
}

func GetResponse(body types.Request) (*http.Response, error) {
    var openAiclient http.Client;
    openAiReq := newOpenAIReq(body);
    return openAiclient.Do(openAiReq);
}



