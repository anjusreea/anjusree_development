// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Employee - Our struct for all employees
type Employee struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	City string `json:"City"`
}

var Employees []Employee

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllEmployees")
	json.NewEncoder(w).Encode(Employees)
	/*dat, err := ioutil.ReadFile("employee.json")
	check(err)
	fmt.Print(string(dat))
	json.NewEncoder(w).Encode(string(dat))*/
	/*test comment*/

}

func returnSingleEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, employee := range Employees {
		if employee.ID == key {
			json.NewEncoder(w).Encode(employee)
		}
	}
}

func createNewEmployee(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Employee struct
	// append this to our Employees array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee Employee
	json.Unmarshal(reqBody, &employee)
	// update our global Employees array to include
	// our new Employee
	Employees = append(Employees, employee)

	json.NewEncoder(w).Encode(employee)
	file, _ := json.MarshalIndent(Employees, "", " ")

	_ = ioutil.WriteFile("employee.json", file, 0644)
}
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, employee := range Employees {
		if employee.ID == id {
			Employees = append(Employees[:index], Employees[index+1:]...)
		}
	}
	// get the body of our POST request
	// unmarshal this into a new Employee struct
	// append this to our Employees array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee Employee
	json.Unmarshal(reqBody, &employee)
	// update our global Employees array to include
	// our new Employee
	Employees = append(Employees, employee)

	json.NewEncoder(w).Encode(employee)
	file, _ := json.MarshalIndent(Employees, "", " ")

	_ = ioutil.WriteFile("employee.json", file, 0644)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, employee := range Employees {
		if employee.ID == id {
			Employees = append(Employees[:index], Employees[index+1:]...)
		}
	}
	file, _ := json.MarshalIndent(Employees, "", " ")

	_ = ioutil.WriteFile("employee.json", file, 0644)

}

func handleRequests() {
	myR := mux.NewRouter().StrictSlash(true)
	myR.HandleFunc("/", homePage)
	myR.HandleFunc("/employees", returnAllEmployees).Methods("GET")
	myR.HandleFunc("/employee", createNewEmployee).Methods("POST")
	myR.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	myR.HandleFunc("/employee/{id}", returnSingleEmployee).Methods("GET")
	myR.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myR))
}

func main() {
	Employees = []Employee{
		Employee{ID: "1", Name: "Emp1", City: "City1"},
		Employee{ID: "2", Name: "Emp2", City: "City2"},
		Employee{ID: "3", Name: "Emp3", City: "City3"},
		Employee{ID: "4", Name: "Emp4", City: "City4"},
		Employee{ID: "5", Name: "Emp5", City: "City5"},
	}
	handleRequests()
}
