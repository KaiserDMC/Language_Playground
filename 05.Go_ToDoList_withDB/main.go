package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id        string
	Label     string
	Date      time.Time
	Completed bool
}

var data = map[string][]Todo{
	"TodoList": {
		{Id: uuid.New().String(), Label: "Template", Date: time.Now(), Completed: false},
		{Id: uuid.New().String(), Label: "Test Number 2", Date: time.Now(), Completed: true},
	},
}

func initDataFromDB(db *sql.DB) error {
	rows, err := db.Query("SELECT id, label, date, completed FROM todos")
	if err != nil {
		return err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Label, &todo.Date, &todo.Completed)
		if err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	// Sort todos by Completed (false first, true later)
	sort.Slice(todos, func(i, j int) bool {
		return !todos[i].Completed && todos[j].Completed
	})

	data["TodoList"] = todos

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		label := r.FormValue("label")
		date := r.FormValue("date")

		newTodo := Todo{
			Id:        uuid.New().String(),
			Label:     label,
			Completed: false,
		}

		// Parse the date
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			http.Error(w, "Error parsing date", http.StatusBadRequest)
			return
		}

		newTodo.Date = parsedDate

		// Open SQLite database
		db, err := sql.Open("sqlite3", "todos.db")
		if err != nil {
			http.Error(w, "Error opening database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Insert the new Todo into the database
		_, err = db.Exec(`INSERT INTO todos (id, label, date, completed) VALUES (?, ?, ?, ?)`, newTodo.Id, newTodo.Label, newTodo.Date, newTodo.Completed)
		if err != nil {
			fmt.Println("Error inserting todo:", err)
			http.Error(w, "Error inserting todo", http.StatusInternalServerError)
			return
		}

		// Append the newTodo to the in-memory data
		data["TodoList"] = append(data["TodoList"], newTodo)

		// // Trigger htmx refresh to update the Todo list
		w.Header().Set("HX-Trigger", "refresh")
		fmt.Println("Refresh triggered")

		for key, value := range data {
			fmt.Printf("Key: %s, Value: %v\n", key, value)
		}

		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	// Check if the SQLite database file exists
	_, err := os.Stat("todos.db")
	if os.IsNotExist(err) {
		// If the database file does not exist, create it
		db, err := sql.Open("sqlite3", "todos.db")
		if err != nil {
			fmt.Println("Error creating database:", err)
			return
		}
		defer db.Close()

		// Create Todo table
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS todos (
				id TEXT PRIMARY KEY,
				label TEXT,
				date DATETIME,
				completed BOOLEAN
			)
		`)
		if err != nil {
			fmt.Println("Error creating table:", err)
			return
		}
	} else if err == nil {
		// If the database file exists, open it and fetch data
		db, err := sql.Open("sqlite3", "todos.db")
		if err != nil {
			fmt.Println("Error opening database:", err)
			return
		}
		defer db.Close()

		// Initialize data from the database
		err = initDataFromDB(db)
		if err != nil {
			fmt.Println("Error initializing data from database:", err)
			return
		}
	} else {
		fmt.Println("Error checking if database file exists:", err)
		return
	}

	// Start HTTP server
	http.HandleFunc("/", handler)
	http.HandleFunc("/addTodo", addTodoHandler)

	fmt.Println("Server is running on :5050...")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
