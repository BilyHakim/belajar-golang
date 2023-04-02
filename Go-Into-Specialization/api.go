package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Bily", Age: 22, Division: "IT"},
	{ID: 2, Name: "Hakim", Age: 20, Division: "Finance"},
	{ID: 3, Name: "Erlangga", Age: 24, Division: "IT"},
}

var PORT2 = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT2)
	http.ListenAndServe(PORT2, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}
