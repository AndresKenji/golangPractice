package api

import "net/http"

func index(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world from my api server!"))
}