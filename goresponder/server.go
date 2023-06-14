package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"database/sql" // sql connection
	_ "github.com/go-sql-driver/mysql" // mysql driver
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
	http.HandleFunc("/set", sendName) // send name here
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
	fmt.Println("GET params were:", r.URL.Query())
	db, err := sql.Open("mysql", "root:password@tcp(10.217.4.44:3306)/names")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sending name insert...")
	log.Println("INSERT INTO names(name) VALUES " + r.URL.Query().Get("name"))
	rows, err := db.Query("INSERT INTO names (name) VALUES ('" + r.URL.Query().Get("name") + "')")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rows)
}