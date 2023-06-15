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
	db *sql.DB // pass handle
}

const (
	PORT = "PORT"
	PROVIDER = "PROVIDER"
)

func main() {
	// open connection
	db, err := sql.Open("mysql", "root:password@tcp(10.217.4.44:3306)/names")
	if err != nil {
		log.Fatal(err)
	}

	a := initApi(db)
	defer db.Close()

	http.HandleFunc("/set", a.sendName) // send name here
	port, set := os.LookupEnv(PORT)
	if !set {
		port = "8080"
	}

	log.Println(fmt.Sprintf("Listening on port: %s", port))
	log.Println(fmt.Sprintf("Configured for: %s", a.provider))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func initApi(db *sql.DB) *api {
	provider, set := os.LookupEnv(PROVIDER)
	if !set {
		provider = "web"
	}

	return &api{
		provider: provider,
		httpClient: &http.Client{Timeout: time.Second * 5},
		db: db, // pass handle
	}
}

func (a *api) sendName(w http.ResponseWriter, r *http.Request) {
	log.Println("Received:", r.URL.Query())
	
	log.Println("Sending name insert...")
	log.Println("INSERT INTO names(name) VALUES " + r.URL.Query().Get("name"))
	_, err := a.db.Query("INSERT INTO names (name) VALUES ('" + r.URL.Query().Get("name") + "')")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Name inserted successfully!")

}