package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
	"time"
)

// CreateLeave collects user input from the form and creates a new leave request
func CreateLeave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	startDate, err := time.Parse("2006-01-02", r.FormValue("startDate"))
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", r.FormValue("endDate"))
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	leave := database.Leave{
		Uuid:         GetNewUuid(),
		EmployeeUuid: r.FormValue("employeeUuid"),
		StartDate:    startDate,
		EndDate:      endDate,
		LeaveType:    r.FormValue("leaveType"),
	}

	database.AddLeave(leave, w, r)

	http.Redirect(w, r, "/leaves?type=success&message=Leave+request+created+successfully!", http.StatusSeeOther)
}

// UpdateLeave collects user input and updates an existing leave request
func UpdateLeave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	leaveUuid := r.FormValue("leaveUuid")
	leave := database.GetLeaveByUuid(leaveUuid, w, r)
	if (leave == database.Leave{}) {
		fmt.Println("Leave request does not exist")
		http.Redirect(w, r, "/leaves?type=error&message=Leave+request+not+found!", http.StatusSeeOther)
		return
	}

	startDate, err := time.Parse("2006-01-02", r.FormValue("startDate"))
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", r.FormValue("endDate"))
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	leave.EmployeeUuid = r.FormValue("employeeUuid")
	leave.StartDate = startDate
	leave.EndDate = endDate
	leave.LeaveType = r.FormValue("leaveType")

	database.UpdateLeaveInfo(leave, w, r)

	http.Redirect(w, r, "/leave/"+leaveUuid+"?type=success&message=Leave+request+updated+successfully!", http.StatusSeeOther)
}

// DeleteLeave deletes an existing leave request
func DeleteLeave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	leaveUuid := r.FormValue("leaveUuid")
	leave := database.GetLeaveByUuid(leaveUuid, w, r)
	if (leave == database.Leave{}) {
		fmt.Println("Leave request does not exist")
		http.Redirect(w, r, "/leaves?type=error&message=Leave+request+not+found!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("User did not confirm deletion")
		http.Redirect(w, r, "/leave/"+leaveUuid+"?type=error&message=Confirm+deletion!", http.StatusSeeOther)
		return
	}

	database.Deleteleave(leaveUuid, w, r)

	http.Redirect(w, r, "/leaves?type=success&message=Leave+request+deleted+successfully!", http.StatusSeeOther)
}
