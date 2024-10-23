package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
)

// CreateEmployee collects user input from the form and creates a new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employee := database.Employee{
		Uuid:            GetNewUuid(),
		Last_name:       r.FormValue("lastName"),
		First_name:      r.FormValue("firstName"),
		Email:           r.FormValue("email"),
		Phone_number:    r.FormValue("phoneNumber"),
		Department_uuid: r.FormValue("departmentId"),
		Position_uuid:   r.FormValue("positionId"),
		Superior_uuid:   r.FormValue("superiorId"),
	}

	database.AddEmployee(employee, w, r)

	http.Redirect(w, r, "/employees?type=success&message=Employee+created+successfully!", http.StatusSeeOther)
}

// UpdateEmployee collects user input and updates an existing employee
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employeeUuid := r.FormValue("employeeUuid")
	employee := database.GetEmployeeByUuid(employeeUuid, w, r)
	if (employee == database.EmployeeInfo{}) {
		fmt.Println("Employee does not exist")
		http.Redirect(w, r, "/employees?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	newEmployee := database.Employee{
		Uuid:            employeeUuid,
		Last_name:       r.FormValue("lastName"),
		First_name:      r.FormValue("firstName"),
		Email:           r.FormValue("email"),
		Phone_number:    r.FormValue("phoneNumber"),
		Department_uuid: r.FormValue("departmentId"),
		Position_uuid:   r.FormValue("positionId"),
		Superior_uuid:   r.FormValue("superiorId"),
	}

	database.UpdateEmployeeInfo(newEmployee, w, r)

	http.Redirect(w, r, "/employee/"+employeeUuid+"?type=success&message=Employee+updated+successfully!", http.StatusSeeOther)
}

// DeleteEmployee deletes an existing employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employeeUuid := r.FormValue("employeeUuid")
	employee := database.GetEmployeeByUuid(employeeUuid, w, r)
	if (employee == database.EmployeeInfo{}) {
		fmt.Println("Employee does not exist")
		http.Redirect(w, r, "/employees?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("User did not confirm deletion")
		http.Redirect(w, r, "/employee/"+employeeUuid+"?type=error&message=Confirm+deletion!", http.StatusSeeOther)
		return
	}

	database.DeleteEmployee(employeeUuid, w, r)

	http.Redirect(w, r, "/employees?type=success&message=Employee+deleted+successfully!", http.StatusSeeOther)
}
