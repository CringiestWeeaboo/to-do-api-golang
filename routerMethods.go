package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tasks, err := db.Exec("SELECT task FROM tasks")
	if err != nil {
		log.Fatalf("Невозможно получить все задачи из таблицы:%v", err)
	}
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := mux.Vars(r)

}
