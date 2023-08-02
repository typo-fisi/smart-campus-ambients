package main

import (
    "ambients-end/src/handlers"
	"net/http"
)

func main() {
    http.HandleFunc("/assignments", handlers.AssignmentsHandler);
    http.ListenAndServe(":4000", nil);
}


