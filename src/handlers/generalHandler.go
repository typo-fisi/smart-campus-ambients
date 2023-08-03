package handlers

import "net/http"

func GeneralHandle(w http.ResponseWriter, r *http.Request) {
    if ((*r).Method == "OPTIONS") {
        EnableCors(&w);
        w.WriteHeader(200);
    }
}
