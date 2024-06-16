package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type handler struct {
	db  *sql.DB
	mux http.ServeMux
}

func main() {
	var (
		router = mux.NewRouter()
		h      = &handler{}
		err    error
	)

	h.db, err = initDb("tasks.db")
	if err != nil {
		log.Fatalf("Невозможно открыть базу данных: %v", err)
	}
	defer h.db.Close()

	err = h.createTable()
	if err != nil {
		log.Fatalf("Невозможно создать таблицу: %v", err)
	}

	router.HandleFunc("/api/tasks", getTasks(h.db)).Methods("GET")
	router.HandleFunc("/api/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/api/createTask", createTask).Methods("POST")
	router.HandleFunc("/api/updateTask/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/api/deleteTask/{id}", deleteTask).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
