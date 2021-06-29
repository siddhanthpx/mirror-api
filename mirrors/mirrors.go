package mirrors

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func FindMirrors() []string {

	url := "https://www.debian.org/mirror/list"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`ftp\.\w\w\.debian\.org/debian/`)
	links := re.FindAllString(string(data), -1)

	defer response.Body.Close()
	return links

}
