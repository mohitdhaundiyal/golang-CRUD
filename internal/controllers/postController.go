package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/md/go-pro-main/internal/config"
	"github.com/md/go-pro-main/internal/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
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
	config.DB.First(&post, postID)
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	config.DB.Create(&post)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get ID from url
	params := mux.Vars(r)
	// get post by ID
	postID := params["id"]
	var post models.Post
	config.DB.First(&post, postID)
	// get updated json data from body
	json.NewDecoder(r.Body).Decode(&post)
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

	json.NewEncoder(w).Encode("post deleted.")

}
