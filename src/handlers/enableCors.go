package handlers

import "net/http"

func EnableCors(w *http.ResponseWriter) {
    (*w).Header().Add("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "*")
    (*w).Header().Set("Access-Control-Allow-Headers", "*")

}
