package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var (
	tasks  = []Task{}
	mutex  = &sync.Mutex{}
	nextID = 1
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func markTaskDone(w http.ResponseWriter, r *http.Request) {
	var id int
	if _, err := fmt.Sscanf(r.URL.Query().Get("id"), "%d", &id); err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("static", "index.html")
	http.ServeFile(w, r, path)
}

func main() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/add", addTask)
	http.HandleFunc("/done", markTaskDone)

	fmt.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
		os.Exit(1)
	}
}
