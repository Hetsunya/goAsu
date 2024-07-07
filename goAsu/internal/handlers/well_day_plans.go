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

// getWellDayPlans возвращает плановые данные по заданной скважине.
// @Summary Получение плановых данных по скважине
// @Description Возвращает плановые данные по заданной скважине
// @Tags well_day_plans
// @Produce json
// @Param well query int true "ID скважины"
// @Success 200 {array} models.WellDayPlan
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /well_day_plans [get]
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

// createWellDayPlan создает новый плановый день для заданной скважины.
// @Summary Создание планового дня
// @Description Создает новый плановый день для заданной скважины
// @Tags well_day_plans
// @Accept json
// @Produce json
// @Param well body models.WellDayPlan true "Создаваемый плановый день"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /well_day_plans [post]
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

// updateWellDayPlan обновляет плановый день для заданной скважины.
// @Summary Обновление планового дня
// @Description Обновляет плановый день для заданной скважины
// @Tags well_day_plans
// @Accept json
// @Produce json
// @Param well body models.WellDayPlan true "Обновляемый плановый день"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /well_day_plans [put]
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

// deleteWellDayPlan удаляет плановый день для заданной скважины.
// @Summary Удаление планового дня
// @Description Удаляет плановый день для заданной скважины
// @Tags well_day_plans
// @Param id query int true "ID планового дня"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /well_day_plans [delete]
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
