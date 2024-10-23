package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
)

/*
! CreateDepartment collects the user input from the corresponding form
! create a department in a database.deartment struct type
! sends it to the database function to store it
! redirects the user to the corresponding page
*/
func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	departmentName := r.FormValue("departmentName")
	department := database.GetDepartmentByName(departmentName, w, r)
	if (department != database.Department{}) {
		fmt.Println("department already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Department+name+already+exist+!", http.StatusSeeOther)
		return
	}

	leaderFirstName := r.FormValue("leaderFirstName")
	leaderLastName := r.FormValue("leaderLastName")
	departmentLeader := database.GetEmployeeByName(leaderFirstName, leaderLastName, w, r)
	if (departmentLeader == database.Employee{}) {
		fmt.Println("Employee does not exist") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	departmentUuid := GetNewUuid()

	newdepartment := database.Department{Uuid: departmentUuid, Department_leader_uuid: departmentLeader.Uuid, Name: departmentName}
	database.AddDepartment(newdepartment, w, r)

	http.Redirect(w, r, "/department/"+departmentUuid+"?type=success&message=Department+created+successfuly+!", http.StatusSeeOther)
}

/*
! UpdateDepartment collects the user input from the corresponding form
! create a department in a database.department struct type
! sends it to the database function to update it
! redirects the user to the corresponding page
*/
func UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	departmentUuid := r.FormValue("departmentUuid")
	department := database.GetDepartmentByUuid(departmentUuid, w, r)
	if (department == database.Department{}) {
		fmt.Println("department does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/?type=error&message=department+not+found+!", http.StatusSeeOther)
		return
	}

	newDepartmentName := r.FormValue("name")
	if newDepartmentName == "" {
		fmt.Println("name empty") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/department/"+departmentUuid+"?type=error&message=Cannot+have+empty+name+!", http.StatusSeeOther)
		return
	}

	leaderFirstName := r.FormValue("leaderFirstName")
	leaderLastName := r.FormValue("leaderLastName")
	departmentLeader := database.GetEmployeeByName(leaderFirstName, leaderLastName, w, r)
	if (departmentLeader == database.Employee{}) {
		fmt.Println("Employee does not exist") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/?type=error&message=Employee+not+found!", http.StatusSeeOther)
		return
	}

	newdepartment := database.Department{Uuid: departmentUuid, Department_leader_uuid: departmentLeader.Uuid, Name: newDepartmentName}
	database.UpdateDepartmentInfo(newdepartment, w, r)

	http.Redirect(w, r, "/department/"+departmentUuid+"?type=success&message=Department+successfully+update+!", http.StatusSeeOther)
}

/*
! DeleteDepartment collects the user input from the corresponding form
! sends it to the database function to delete it
! redirects the user to the corresponding page
*/
func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	departmentUuid := r.FormValue("departmentUuid")
	department := database.GetDepartmentByUuid(departmentUuid, w, r)
	if (department == database.Department{}) {
		fmt.Println("department does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/?type=error&message=department+not+found+!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/department/"+departmentUuid+"?type=error&message=Confirm+deletion+!", http.StatusSeeOther)
		return
	} else {
		database.DeleteDepartment(departmentUuid, w, r)
	}

	//Send confirmation message
	http.Redirect(w, r, "/?type=success&message=Department+deleted+!", http.StatusSeeOther)
}
