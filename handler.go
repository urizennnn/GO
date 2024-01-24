package main

import "net/http"

func handlerReadines(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
