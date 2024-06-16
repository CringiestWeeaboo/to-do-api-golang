package main

import (
	"database/sql"
	"fmt"
	"github.com/goodsign/monday"
	"log"
	"time"
)

// TODO: вынести запросы в константы и к ним уже обращаться в методах
type DbMethods interface {
	createItem(task string)
	updateItem(ID int, newTask string)
	deleteItem(ID int)
}

type tasksDatabase struct {
	db *sql.DB
}

func NewTasksDatabase(db *sql.DB) DbMethods {
	return tasksDatabase{db: db}
}

func (database tasksDatabase) createItem(task string) {
	date := monday.Format(time.Now(), monday.DefaultFormatRuRUDateTime, monday.LocaleRuRU)
	_, err := database.db.Exec("INSERT INTO tasks (task, createTime) VALUES (?, ?)", task, date)
	if err != nil {
		log.Fatalf("Невозможно записать элемент в таблицу: %v", err)
	}
}

func (database tasksDatabase) updateItem(id int, newTask string) {
	_, err := database.db.Exec("UPDATE tasks SET task = ? WHERE ID = ?", newTask, id)
	if err != nil {
		log.Fatalf("Невозможно обновить элемент таблицы: %v", err)
	}
}

func (database tasksDatabase) deleteItem(id int) {
	_, err := database.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Невозможно удалить элемент из таблицы: %v", err)
	}
}

func initDb(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Невозможно открыть базу данных: %v", err)
	}

	return db, nil
}

func (h handler) createTable() error {
	_, err := h.db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, createTime TEXT)")
	if err != nil {
		return fmt.Errorf("Невозможно создать таблицу базы данных: %v", err)
	}

	return nil
}
