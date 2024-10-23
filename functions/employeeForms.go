package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
	"github.com/google/uuid"
)

// CreateEmployee collects user input from the form and creates a new employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	employee := database.Employee{
		Uuid:          uuid.New().String(),
		Last_name:     r.FormValue("lastName"),
		First_name:    r.FormValue("firstName"),
		Email:         r.FormValue("email"),
		Phone_number:  r.FormValue("phoneNumber"),
		Department_id: r.FormValue("departmentId"),
		Position_id:   r.FormValue("positionId"),
		Superior_id:   r.FormValue("superiorId"),
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
	if (employee == database.Employee{}) {
		fmt.Println("Employee does not exist")
		http.Redirect(w, r, "/employees?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	employee.Last_name = r.FormValue("lastName")
	employee.First_name = r.FormValue("firstName")
	employee.Email = r.FormValue("email")
	employee.Phone_number = r.FormValue("phoneNumber")
	employee.Department_id = r.FormValue("departmentId")
	employee.Position_id = r.FormValue("positionId")
	employee.Superior_id = r.FormValue("superiorId")

	database.UpdateEmployeeInfo(employee, w, r)

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
	if (employee == database.Employee{}) {
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