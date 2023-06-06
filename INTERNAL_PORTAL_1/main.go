package main

import (
	"log"
	"net/http"

	"github.com/gocql/gocql"

	api "nswe.com/events/API"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "my_app"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra:", err)
	}
	defer session.Close()

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")

		api.GetProfileHandler(w, r, session, name)
	})

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
