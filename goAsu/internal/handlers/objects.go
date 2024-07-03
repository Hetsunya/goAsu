package handlers

import (
	"database/sql"
	"encoding/json"
	"goAsu/internal/models"
	"net/http"
	"strconv"
)

func ObjectsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getObjects(db, w, r)
		case "POST":
			createObject(db, w, r)
		case "PUT":
			updateObject(db, w, r)
		case "DELETE":
			deleteObject(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getObjects(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, type FROM objects")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var objects []models.Object
	for rows.Next() {
		var obj models.Object
		if err := rows.Scan(&obj.ID, &obj.Name, &obj.Type); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		objects = append(objects, obj)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)
}

func createObject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO objects (name, type) VALUES ($1, $2) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, obj.Name, obj.Type).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	obj.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func updateObject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `UPDATE objects SET name=$1, type=$2 WHERE id=$3`
	res, err := db.Exec(sqlStatement, obj.Name, obj.Type, obj.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "No rows affected", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func deleteObject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	sqlStatement := `DELETE FROM objects WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "No rows affected", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
