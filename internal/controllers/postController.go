package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/md/go-pro-main/internal/config"
	"github.com/md/go-pro-main/internal/models"
)

var message = make(map[string]string)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GoLang CRUD API.")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post []models.Post
	config.DB.Find(&post)
	json.NewEncoder(w).Encode(post)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post []models.Post
	postID := params["id"]
	res := config.DB.First(&post, postID)

	if res.RowsAffected == 0 {
		message["message"] = "Post does not exists."
		json.NewEncoder(w).Encode(message)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	// validations
	if post.Title == "" {
		message["message"] = "Title required."
		json.NewEncoder(w).Encode(message)
		return
	} else if post.Description == "" {
		message["message"] = "Description required."
		json.NewEncoder(w).Encode(message)
		return
	}

	// create post
	config.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get ID from url
	params := mux.Vars(r)
	postID := params["id"]

	var post models.Post
	// get post from ID
	config.DB.First(&post, postID)

	// get updated json data from body
	json.NewDecoder(r.Body).Decode(&post)

	// validations
	if post.Title == "" {
		message["message"] = "Title required."
		json.NewEncoder(w).Encode(message)
		return
	} else if post.Description == "" {
		message["message"] = "Description required."
		json.NewEncoder(w).Encode(message)
		return
	}

	// fmt.Print(r.Body)
	// save updated value to DB
	config.DB.Save(&post)
	json.NewEncoder(w).Encode(&post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	postID := params["id"]

	var post models.Post
	config.DB.Delete(&post, postID)

	message["message"] = "Post deleted."
	json.NewEncoder(w).Encode(message)

}
