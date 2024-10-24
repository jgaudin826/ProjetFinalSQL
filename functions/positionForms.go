package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
	"strconv"
)

/*
! CreatePosition collects the user input from the corresponding form
! create a position in a database.deartment struct type
! sends it to the database function to store it
! redirects the user to the corresponding page
*/
func CreatePosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	positionName := r.FormValue("positionName")
	position := database.GetPositionByName(positionName, w, r)
	if (position != database.Position{}) {
		fmt.Println("position already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Position+name+already+exist+!", http.StatusSeeOther)
		return
	}

	salary := r.FormValue("salary")
	salaryInt, _ := strconv.Atoi(salary)

	positionUuid := GetNewUuid()

	newposition := database.Position{Uuid: positionUuid, Title: positionName, Salary: salaryInt}
	database.AddPosition(newposition, w, r)

	http.Redirect(w, r, "/position/"+positionUuid+"?type=success&message=Position+created+successfuly+!", http.StatusSeeOther)
}

/*
! UpdatePosition collects the user input from the corresponding form
! create a position in a database.position struct type
! sends it to the database function to update it
! redirects the user to the corresponding page
*/
func UpdatePosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	positionUuid := r.FormValue("positionUuid")
	position := database.GetPositionByUuid(positionUuid, w, r)
	if (position == database.Position{}) {
		fmt.Println("position does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/?type=error&message=Position+not+found+!", http.StatusSeeOther)
		return
	}

	newpositionName := r.FormValue("name")
	if newpositionName == "" {
		fmt.Println("name empty") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/position/"+positionUuid+"?type=error&message=Cannot+have+empty+name+!", http.StatusSeeOther)
		return
	}

	salary := r.FormValue("salary")
	salaryInt, _ := strconv.Atoi(salary)

	newposition := database.Position{Uuid: positionUuid, Title: newpositionName, Salary: salaryInt}
	database.UpdatePositionInfo(newposition, w, r)

	http.Redirect(w, r, "/position/"+positionUuid+"?type=success&message=Position+successfully+update+!", http.StatusSeeOther)
}

/*
! Deleteposition collects the user input from the corresponding form
! sends it to the database function to delete it
! redirects the user to the corresponding page
*/
func DeletePosition(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	positionUuid := r.FormValue("positionUuid")
	position := database.GetPositionByUuid(positionUuid, w, r)
	if (position == database.Position{}) {
		fmt.Println("position does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/?type=error&message=Position+not+found+!", http.StatusSeeOther)
		return
	}

	database.DeletePosition(positionUuid, w, r)

	//Send confirmation message
	http.Redirect(w, r, "/?type=success&message=Position+deleted+!", http.StatusSeeOther)
}
