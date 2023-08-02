package main

import (
	"ambients-end/src/handlers"
	"net/http"
)

func main() {
    http.HandleFunc("/api/assignments/", handlers.AssignmentsHandler);
    http.HandleFunc("/api/ambients/", handlers.AmbientsHandler);
    http.ListenAndServe(":4000", nil);
}


