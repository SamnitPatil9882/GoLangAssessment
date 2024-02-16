package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type showWebsites struct {
	Names []string `json:"websites"`
}

var map1 map[string]string
var wg sync.WaitGroup

func main() {
	if map1 == nil {
		map1 = make(map[string]string)
	}
	fmt.Println("Starting Server...")
	r := mux.NewRouter()

	r.HandleFunc("/websites", WebsiteHandler).Methods(http.MethodPost)
	r.HandleFunc("/websites", WebsiteStatuses).Methods(http.MethodGet)
	r.HandleFunc("/website", showStatus).Methods(http.MethodGet)

	http.ListenAndServe("localhost:3000", r)
	fmt.Println("Server started !!")
}

func WebsiteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In handler=> POST/websites")

	webName := showWebsites{}
	err := json.NewDecoder(r.Body).Decode(&webName)
	if err != nil {
		fmt.Println(fmt.Errorf("error !! Something went wrong"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(webName.Names) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Empty list of websites provided.")
		return
	}

	for _, url := range webName.Names {
		if !isValidURL(url) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid URL: %s", url)
			return
		}
	}

	wg.Add(len(webName.Names))
	for _, url := range webName.Names {
		fmt.Println("Your Url : ", url)
		go checkStatus(url)
	}

	fmt.Fprint(w, "POST/website Completed !!")
}

func WebsiteStatuses(w http.ResponseWriter, r *http.Request) {
	log.Println("In handler=> GET/websites")

	fmt.Fprint(w, "Your Access Values : ", map1)
}

func showStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	m1 := r.Form
	val, ok := m1["name"]
	if !ok || len(val) == 0 || !isValidURL(val[0]) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid or missing URL parameter.")
		return
	}

	res, exists := map1[val[0]]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v is now Down !!", val[0])
		return
	}
	fmt.Fprintf(w, "%v is now %v !!", val[0], res)
}

func checkStatus(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("error !! Something went wrong !!")
		return
	}

	if code := res.StatusCode; code >= 200 && code < 300 {
		map1[url] = "UP"
	} else {
		map1[url] = "DOWN"
	}

	wg.Done()
}

func checkStatusForever(url string) {
	for {
		res, err := http.Get(url)

		if err != nil {
			map1[url] = "DOWN"
			return
		}

		if code := res.StatusCode; code >= 200 && code < 300 {
			map1[url] = "UP"
		}

		time.Sleep(time.Minute)
	}
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
