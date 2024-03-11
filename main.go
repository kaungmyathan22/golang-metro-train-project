package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/kaungmyathan22/golang-metro-train/api"
	"github.com/kaungmyathan22/golang-metro-train/database"

	_ "github.com/mattn/go-sqlite3"
)

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

func main() {
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// Create tables
	database.Initialize(db)
	api.DB = db
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := api.TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
