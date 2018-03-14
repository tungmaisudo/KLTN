package main

import (
	"KLTN/data-stream-backend/Cassandra"
	"KLTN/data-stream-backend/Stream"
	"KLTN/data-stream-backend/WifiCassandra"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	err := Stream.Connect(
		"rvkndq55xkj5", //key
		"ymmbngy5vncmksj4wv5e8fkcp84rv93ajzcu2hbxzfhwgy5nuyhm5dk3bk9fzd7p", //secret
		"us-east", //region
		"33835")   //app id
	if err != nil {
		log.Fatal("Could not connect to Stream, abort")
	} else {
		log.Printf("Success Connect to Stream")
	}

	// url := "https://data.cityofnewyork.us/api/id/varh-9tsp.json"
	// dataList, err := Data.GetDataByUrl(url)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	log.Println("Connect cityofnewyork successlly")
	// }

	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()
	// for index := range dataList {
	// 	insertWifiToCassandra(dataList[index])
	// }
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get-all-wifi", WifiCassandra.GetAllWifi).Methods("Get")
	router.HandleFunc("/post-all-wifi", WifiCassandra.PostWifiData).Methods("POST")
	router.HandleFunc("/delete-all-wifi", WifiCassandra.DeleteWifiData).Methods("DELETE")

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	handler := c.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	log.Fatal(srv.ListenAndServe())

}
