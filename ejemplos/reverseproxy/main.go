package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	demoURL, err := url.Parse("https://otherurl.fake.com")
	if err != nil {
		log.Panic(err.Error())
	}
	proxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = demoURL.Host
		r.URL.Host = demoURL.Host
		r.URL.Scheme = demoURL.Scheme
		r.RequestURI = ""
		response, err := http.DefaultClient.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}
		for key, values := range response.Header {
			for _, value := range values {
				w.Header().Set(key, value)
			}
		}

		w.WriteHeader(response.StatusCode)
		io.Copy(w, response.Body)

	})
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", proxy)
}
