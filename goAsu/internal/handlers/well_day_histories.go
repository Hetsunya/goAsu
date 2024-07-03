package handlers

import (
	"database/sql"
	"encoding/json"
	"goAsu/internal/models"
	"net/http"
	"strconv"
)

func WellDayHistoriesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getWellDayHistories(db, w, r)
		case "POST":
			createWellDayHistory(db, w, r)
		case "PUT":
			updateWellDayHistory(db, w, r)
		case "DELETE":
			deleteWellDayHistory(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getWellDayHistories(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT well, date_fact, debit, ee_consume, expenses, pump_operating FROM well_day_histories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var histories []models.WellDayHistory
	for rows.Next() {
		var history models.WellDayHistory
		if err := rows.Scan(&history.Well, &history.DateFact, &history.Debit, &history.EEConsume, &history.Expenses, &history.PumpOperating); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		histories = append(histories, history)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(histories)
}

func createWellDayHistory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var history models.WellDayHistory
	if err := json.NewDecoder(r.Body).Decode(&history); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO well_day_histories (well, date_fact, debit, ee_consume, expenses, pump_operating) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, history.Well, history.DateFact, history.Debit, history.EEConsume, history.Expenses, history.PumpOperating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func updateWellDayHistory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var history models.WellDayHistory
	if err := json.NewDecoder(r.Body).Decode(&history); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE well_day_histories SET debit=$1, ee_consume=$2, expenses=$3, pump_operating=$4 WHERE well=$5 AND date_fact=$6`
	_, err := db.Exec(sqlStatement, history.Debit, history.EEConsume, history.Expenses, history.PumpOperating, history.Well, history.DateFact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func deleteWellDayHistory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	well, err := strconv.Atoi(r.URL.Query().Get("well"))
	if err != nil {
		http.Error(w, "Invalid Well ID", http.StatusBadRequest)
		return
	}
	dateFact := r.URL.Query().Get("date_fact")
	if dateFact == "" {
		http.Error(w, "Invalid Date", http.StatusBadRequest)
		return
	}

	sqlStatement := `DELETE FROM well_day_histories WHERE well=$1 AND date_fact=$2`
	_, err = db.Exec(sqlStatement, well, dateFact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
