// employee.go
package main

import (
	"ProjetFinalSQL/database"
	"html/template"
	"net/http"
	"log"
)

type EmployeeInfo struct {
	Uuid          string
	First_name    string
	Last_name     string
	Email         string
	Phone_number  string
	Department_id string
	Department_name string 
	Position_id   string
	Position_name   string 
	Superior_id   string
	Superior_name   string 
}

type Event struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func Employee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employeeUuid := r.URL.Query().Get("uuid")

	log.Printf("Employee UUID: %v", employeeUuid)

	employeeInfo := database.GetEmployeeByUuid(employeeUuid, w, r) 
	if (employeeInfo == database.EmployeeInfo{}) { 
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	log.Printf("Employee Info: %v", employeeInfo)

	leaves := database.GetLeaveByEmployeeUuid(employeeUuid, w, r)
	
	var events []Event
	for _, leave := range leaves {
		events = append(events, Event{
			Title: leave.LeaveType,
			Start: leave.StartDate.Format("2006-01-02"),
			End:   leave.EndDate.Format("2006-01-02"),   
		})
	}

	log.Print("Events JSON: ", events)

	tmpl, err := template.ParseFiles("./templates/employee.html") 
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, struct {
		Employee EmployeeInfo
		Events   []Event
	}{
		Employee: EmployeeInfo{
			Uuid:          employeeInfo.Uuid, 
			First_name:    employeeInfo.First_name,
			Last_name:     employeeInfo.Last_name,
			Email:         employeeInfo.Email,
			Phone_number:  employeeInfo.Phone_number,
			Department_id: employeeInfo.Department_uuid,
			Position_id:   employeeInfo.Position_uuid,
			Superior_id:   employeeInfo.Superior_uuid,
			Department_name: employeeInfo.Department_name,
			Position_name:   employeeInfo.Position_name,
			Superior_name:   employeeInfo.Superior_name,  
		},
		Events: events,
	})
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}