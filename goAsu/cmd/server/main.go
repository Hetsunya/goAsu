package main

import (
	"goAsu/internal/database"
	"goAsu/internal/handlers"
	"goAsu/internal/models"
	"log"
	"net/http"

	_ "goAsu/docs"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title NeftDobicha API
// @version 1.0
// @description REST API для обработки данных, хранимых в СУБД предприятия "НефтьДобыча".

// @host localhost:8080
// @BasePath /
func main() {
	db := database.InitDB(models.BASE_IP,
		models.PORT,
		models.USERNAME,
		models.PASSWORD,
		models.BASENAME)
	defer db.Close()

	http.HandleFunc("/objects", handlers.ObjectsHandler(db))
	http.HandleFunc("/wells", handlers.WellsHandler(db))
	http.HandleFunc("/well_day_histories", handlers.WellDayHistoriesHandler(db))
	http.HandleFunc("/well_day_plans", handlers.WellDayPlansHandler(db))

	// Добавьте маршрут для Swagger UI
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
