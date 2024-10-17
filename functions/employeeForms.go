package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
! CreatePost collects the user input from the corresponding form
! Check if user is connected and has the rigth to do that action
! Saves the inputed files to the corresponding folders and renames them
! create a post in a database.post struct type
! sends it to the database function to create it
! redirects the user to the corresponding page
*/
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Check if user connected
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	}

	postUuid := GetNewUuid()
	title := r.FormValue("title")
	content := r.FormValue("content")
	id := r.FormValue("communityId")
	communityId, _ := strconv.Atoi(id)

	r.ParseMultipartForm(10 << 20)

	// Get image or video file or link from user
	mediaPath := "/static/images/mediaTemplate.png"
	mediaType := ""

	profileOption := r.FormValue("mediaOption")
	if profileOption == "link" {
		mediaPath = r.FormValue("mediaLink")
	} else {
		profile, handler, err := r.FormFile("media")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			mediaPath = "/static/images/mediaTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" {
				mediaPath = "/static/images/posts/" + postUuid + ext
				mediaType = "image"
				GetFileFromForm(profile, handler, err, mediaPath)
			} else if e == ".mp4" || e == ".webm" || e == ".ogg" {
				mediaPath = "/static/images/posts/" + postUuid + ext
				mediaType = "video"
				GetFileFromForm(profile, handler, err, mediaPath)
			} else {
				fmt.Println("The file is  not in an image or video format")
				return //if not an image or video format
			}
		}
	}
	post := database.Post{Id: 0, Uuid: postUuid, Title: title, Content: content, Media: mediaPath, MediaType: mediaType, User_id: user.Id, Community_id: communityId, Created: (time.Now())}
	database.AddPost(post, w, r)
	createdPost := database.GetPostByUuid(postUuid, w, r)
	like := database.Like{Id: 0, Rating: "like", Comment_id: 0, Post_id: createdPost.Id, User_id: user.Id}
	database.AddLike(like, w, r)

	http.Redirect(w, r, "/post/"+postUuid+"?type=success&message=Post+successfully+created+!", http.StatusSeeOther)
}

/*
! UpdatePost collects the user input from the corresponding form
! Check if user is connected and has the rigth to do that action
! Saves the inputed files to the corresponding folders and renames them / deletes the previous ones if they are modified
! create a post in a database.post struct type
! sends it to the database function to update it
! redirects the user to the corresponding page
*/
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	postId := r.FormValue("postId")
	id, _ := strconv.Atoi(postId)
	post := database.GetPostById(id, w, r)
	if (post == database.PostInfo{}) {
		fmt.Println("post does not exist") // TO-DO : send error post not found
		http.Redirect(w, r, "/?type=error&message=Post+not+found+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	} else if post.User_id != user.Id {
		fmt.Println("user not author of post") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=User+not+alowed+to+do+this+action+!", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	r.ParseMultipartForm(10 << 20)

	// Get image or video file or link from user
	mediaPath := "/static/images/mediaTemplate.png"
	mediaType := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		DeleteFile(post.Media)
		mediaPath = "/static/images/mediaTemplate.png"
	} else if profileOption == "keep" {
		mediaPath = post.Media
	} else if profileOption == "link" {
		DeleteFile(post.Media)
		mediaPath = r.FormValue("profileLink")
	} else {
		DeleteFile(post.Media)
		profile, handler, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			mediaPath = "/static/images/mediaTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" {
				mediaPath = "/static/images/posts/" + post.Uuid + ext
				mediaType = "image"
				GetFileFromForm(profile, handler, err, mediaPath)
			} else if e == ".mp4" || e == ".webm" || e == ".ogg" {
				mediaPath = "/static/images/posts/" + post.Uuid + ext
				mediaType = "video"
				GetFileFromForm(profile, handler, err, mediaPath)
			} else {
				fmt.Println("The file is  not in an image or video format")
				return //if not an image or video format
			}
		}
	}

	updatedPost := database.Post{Id: post.Id, Uuid: post.Uuid, Title: title, Content: content, Media: mediaPath, MediaType: mediaType, User_id: post.User_id, Community_id: post.Community_id, Created: post.Created}
	database.UpdatePostInfo(updatedPost, w, r)

	http.Redirect(w, r, "/post/"+post.Uuid+"?type=success&message=Post+successfully+update+!", http.StatusSeeOther)
}

/*
! DeletePost collects the user input from the corresponding form
! Check if user is connected and has the rigth to do that action
! Deletes the saved images linked to that community
! sends it to the database function to delete it
! redirects the user to the corresponding page
*/
func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	postId := r.FormValue("postId")
	id, _ := strconv.Atoi(postId)
	post := database.GetPostById(id, w, r)
	if (post == database.PostInfo{}) {
		fmt.Println("post does not exist") // TO-DO : send error post not found
		http.Redirect(w, r, "/?type=error&message=Post+not+found+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	} else if post.User_id != user.Id {
		fmt.Println("user not author of post") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=User+not+alowed+to+do+this+action+!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/post/"+post.Uuid+"?type=error&message=Confim+deletion+!", http.StatusSeeOther)
		return
	} else {
		DeleteFile(post.Media)
		database.DeletePost(post.Id, w, r)
	}

	//Send confirmation message
	http.Redirect(w, r, "/?type=success&message=Post+succesfully+deleted+post+!", http.StatusSeeOther)
}
