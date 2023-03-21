package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func Routes(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Task Manager!"))
	})

	r.Get("/tasks", func(w http.ResponseWriter, r *http.Request) {
		var tasks []Task
		if err := db.Select(&tasks, "SELECT * FROM tasks ORDER BY created_at DESC"); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	r.Post("/tasks", func(w http.ResponseWriter, r *http.Request) {
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		_, err := db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	})

	r.Put("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		_, err := db.Exec("UPDATE tasks SET title=?, description=?, completed=? WHERE id=?", task.Title, task.Description, task.Completed, id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	r.Delete("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	return r
}
