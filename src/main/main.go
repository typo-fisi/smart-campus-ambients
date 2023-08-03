package main

import (
	"ambients-end/src/handlers"
	"net/http"
)

func main() {
    http.HandleFunc("/api/assignments/", handlers.AssignmentsHandler);
    http.HandleFunc("/api/ambients/", handlers.AmbientsHandler);
    http.HandleFunc("/api/professors/", handlers.ProfessorHandler);
    http.ListenAndServe(":4000", nil);
}


