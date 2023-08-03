package handlers

import (
	"net/http"
	"path/filepath"
)

func ProfessorHandler(w http.ResponseWriter, r *http.Request) {
    urlPathString := filepath.Clean((*r).URL.EscapedPath());
    //urlPathElements := strings.Split(urlPathString, "/");

    w.Write([]byte(urlPathString));
}








