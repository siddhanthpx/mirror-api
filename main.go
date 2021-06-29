package main

import (
	"fmt"
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

	for _, link := range mirrorList {

		resp, err := http.Get("http://" + link + "README.html")
		errorHandler("Couldn't GET mirror", err)
		fmt.Println(resp.Status)

	}
}

func errorHandler(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func findFastest(urls []string) {

	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)
}
