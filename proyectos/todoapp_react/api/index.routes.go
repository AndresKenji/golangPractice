package api

import (
	"errors"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from my api server!"))
}

func CheckCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("userid")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	// Echo out the cookie value in the response body.
	w.Write([]byte(cookie.Value))

}
