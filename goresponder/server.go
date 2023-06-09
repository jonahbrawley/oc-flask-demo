package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type api struct {
	provider string
	httpClient *http.Client
	//alertHandler http.HandleFunc // added
}

const (
	PORT = "PORT"
	PROVIDER = "PROVIDER"
)

func main() {
	a := initApi()
	http.HandleFunc("/name", sendName)
	port, set := os.LookupEnv(PORT)
	if !set {
		port = "8080"
	}
	//log.Println("test")
	log.Println(fmt.Sprintf("Listening on port: %s", port))
	log.Println(fmt.Sprintf("Configured for: %s", a.provider))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func initApi() *api {
	provider, set := os.LookupEnv(PROVIDER)
	if !set {
		provider = "web"
	}
	return &api{
		provider: provider,
		httpClient: &http.Client{Timeout: time.Second * 5},
		//alertHandler: sendName,
	}
}

func sendName(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending name...")
}