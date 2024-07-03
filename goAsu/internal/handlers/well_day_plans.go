package handlers

import (
	"database/sql"
	"encoding/json"
	"goAsu/internal/models"
	"net/http"
	"strconv"
)

func WellDayPlansHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getWellDayPlans(db, w, r)
		case "POST":
			createWellDayPlan(db, w, r)
		case "PUT":
			updateWellDayPlan(db, w, r)
		case "DELETE":
			deleteWellDayPlan(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getWellDayPlans(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT well, date_plan, debit, ee_consume, expenses, pump_operating FROM well_day_plans")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var plans []models.WellDayPlan
	for rows.Next() {
		var plan models.WellDayPlan
		if err := rows.Scan(&plan.Well, &plan.DatePlan, &plan.Debit, &plan.EEConsume, &plan.Expenses, &plan.PumpOperating); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		plans = append(plans, plan)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

func createWellDayPlan(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var plan models.WellDayPlan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO well_day_plans (well, date_plan, debit, ee_consume, expenses, pump_operating) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, plan.Well, plan.DatePlan, plan.Debit, plan.EEConsume, plan.Expenses, plan.PumpOperating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

func updateWellDayPlan(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var plan models.WellDayPlan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE well_day_plans SET debit=$1, ee_consume=$2, expenses=$3, pump_operating=$4 WHERE well=$5 AND date_plan=$6`
	_, err := db.Exec(sqlStatement, plan.Debit, plan.EEConsume, plan.Expenses, plan.PumpOperating, plan.Well, plan.DatePlan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

func deleteWellDayPlan(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	well, err := strconv.Atoi(r.URL.Query().Get("well"))
	if err != nil {
		http.Error(w, "Invalid Well ID", http.StatusBadRequest)
		return
	}
	datePlan := r.URL.Query().Get("date_plan")
	if datePlan == "" {
		http.Error(w, "Invalid Date", http.StatusBadRequest)
		return
	}

	sqlStatement := `DELETE FROM well_day_plans WHERE well=$1 AND date_plan=$2`
	_, err = db.Exec(sqlStatement, well, datePlan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
