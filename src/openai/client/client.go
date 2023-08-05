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
    openAiPath := "https://api.openai.com/v1/completions";
    openAiKey := "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";

    bodyStr := new(strings.Builder);
    encodeErr := json.NewEncoder(bodyStr).Encode(body);
    if (encodeErr != nil) { log.Fatal(encodeErr); }

    openAiReq, reqErr := http.NewRequest(http.MethodPost, openAiPath, bytes.NewReader([]byte(bodyStr.String())));
    if (reqErr != nil) { log.Fatal(reqErr); }

    (*openAiReq).Header.Set("Authorization", "Bearer " + openAiKey);
    (*openAiReq).Header.Set("Content-Type", "application/json" + openAiKey);

    return openAiReq;
}

func GetResponse(body types.Request) (*http.Response, error) {
    var openAiclient http.Client;
    openAiReq := newOpenAIReq(body);
    return openAiclient.Do(openAiReq);
}



