package main


import (
	"net/http"
)

func handleError(w http.ResponseWriter, err *http.Request){
	respondWithError(w, 400, "Internal Server Error")
}