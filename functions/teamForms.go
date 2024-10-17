package functions

import (
	"ProjetFinalSQL/database"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

/*
! CreateComment collects the user input from the corresponding form
! Check if user is connected and other potential errors
! create a comment in a database.comment struct type
! sends it to the database function to store it
! redirects the user to the corresponding page
*/
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentid := r.FormValue("commentId")
	commentId, _ := strconv.Atoi(commentid)
	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	content := r.FormValue("content")

	team := database.Team{}
	database.AddTeam(team, w, r)

	http.Redirect(w, r, "/post/"+postUuid+"?type=success&message=Comment+posted+!", http.StatusSeeOther)
}

/*
! UpdateComment collects the user input from the corresponding form
! Check if user is connected and has the rigth to do that action
! create a comment in a database.comment struct type
! sends it to the database function to update it
! redirects the user to the corresponding page
*/
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentId := r.FormValue("commentId")
	commentid, _ := strconv.Atoi(commentId)
	comment := database.GetTeamByUuid(commentid, w, r)
	if (comment == database.CommentInfo{}) {
		fmt.Println("comment does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/search/?type=error&message=Comment+not+found+!", http.StatusSeeOther)
		return
	}

	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	content := r.FormValue("content")

	commentUpdate := database.Comment{Id: 0, User_id: user.Id, Post_id: postId, Comment_id: commentid, Content: content, Created: time.Now()}
	database.AddComment(commentUpdate, w, r)

	http.Redirect(w, r, "/comment/"+commentId+"?type=success&message=Comment+successfully+update+!", http.StatusSeeOther)
}

/*
! DeleteComment collects the user input from the corresponding form
! Check if user is connected and has the rigth to do that action
! sends it to the database function to delete it
! redirects the user to the corresponding page
*/
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentId := r.FormValue("commentId")
	commentid, _ := strconv.Atoi(commentId)
	comment := database.GetCommentById(commentid, w, r)
	if (comment == database.CommentInfo{}) {
		fmt.Println("comment does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=Comment+not+found+!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=Confirm+deletion+!", http.StatusSeeOther)
		return
	} else {
		database.DeleteComment(comment.Id, w, r)
	}

	//Send confirmation message
	http.Redirect(w, r, "/?type=success&message=Comment+deleted+!", http.StatusSeeOther)
}
