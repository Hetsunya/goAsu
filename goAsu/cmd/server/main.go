package main

import (
	"goAsu/internal/database"
	"goAsu/internal/handlers"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db := database.InitDB("109.120.183.88", 5432, "hetsu", "Admin1234567890!", "PostgreSQL-vitalick113")
	defer db.Close()

	http.HandleFunc("/objects", handlers.ObjectsHandler(db))
	http.HandleFunc("/wells", handlers.WellsHandler(db))
	http.HandleFunc("/well_day_histories", handlers.WellDayHistoriesHandler(db))
	http.HandleFunc("/well_day_plans", handlers.WellDayPlansHandler(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
