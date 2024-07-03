package handlers

import (
	"database/sql"
	"encoding/json"
	"goAsu/internal/models"
	"net/http"
	"strconv"
)

func WellsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getWells(db, w, r)
		case "POST":
			createWell(db, w, r)
		case "PUT":
			updateWell(db, w, r)
		case "DELETE":
			deleteWell(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getWells(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT well, ngdu, cdng, kust, mest FROM wells")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var wells []models.Well
	for rows.Next() {
		var well models.Well
		if err := rows.Scan(&well.Well, &well.NGDU, &well.CDNG, &well.Kust, &well.Mest); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		wells = append(wells, well)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wells)
}

func createWell(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var well models.Well
	if err := json.NewDecoder(r.Body).Decode(&well); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO wells (well, ngdu, cdng, kust, mest) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, well.Well, well.NGDU, well.CDNG, well.Kust, well.Mest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(well)
}

func updateWell(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var well models.Well
	if err := json.NewDecoder(r.Body).Decode(&well); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE wells SET ngdu=$1, cdng=$2, kust=$3, mest=$4 WHERE well=$5`
	_, err := db.Exec(sqlStatement, well.NGDU, well.CDNG, well.Kust, well.Mest, well.Well)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(well)
}

func deleteWell(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	well, err := strconv.Atoi(r.URL.Query().Get("well"))
	if err != nil {
		http.Error(w, "Invalid Well ID", http.StatusBadRequest)
		return
	}

	sqlStatement := `DELETE FROM wells WHERE well=$1`
	_, err = db.Exec(sqlStatement, well)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
