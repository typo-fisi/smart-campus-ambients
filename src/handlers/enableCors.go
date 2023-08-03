package handlers

import "net/http"

func EnableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Methods", "\"POST\", \"GET\", \"OPTIONS\", \"PUT\"");
    (*w).Header().Set("Access-Control-Allow-Origin", "*");
    (*w).Header().Set("Access-Control-Allow-Headers", "*");
}
