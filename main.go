package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/siddhanthpx/mirror-api/mirrors"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {

	mirrorList := mirrors.FindMirrors()

	http.HandleFunc("/fastest-mirror", func(rw http.ResponseWriter, r *http.Request) {
		respObject := findFastest(mirrorList)

		respJSON, err := json.Marshal(&respObject)
		errorHandler("Unable to parse response object", err)

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(respJSON)
	})

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server on localhost:8080")
	log.Fatal(srv.ListenAndServe())

}

func findFastest(urls []string) response {

	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		link := url
		go func() {
			start := time.Now()
			_, err := http.Get("http://" + link + "README.html")
			latency := time.Now().Sub(start) / time.Millisecond

			if err == nil {
				urlChan <- link
				latencyChan <- latency
			}

		}()
	}

	return response{<-urlChan, <-latencyChan}
}

func errorHandler(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
