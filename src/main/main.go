package main

import (
	"ambients-end/src/handlers"
	"log"
	"net/http"
)

func main() {
    port := "4000";

    http.HandleFunc("/", handlers.GeneralHandle);
    http.HandleFunc("/api/assignments/", handlers.AssignmentsHandler);
    http.HandleFunc("/api/ambients/", handlers.AmbientsHandler);
    http.HandleFunc("/api/helper", handlers.Helper);

    log.Println("Attempting to listen and serve on port " + port + "...");

    startErr := http.ListenAndServe(":" + port, nil);
    if (startErr != nil) {
        log.Fatal(startErr);
    }
}
