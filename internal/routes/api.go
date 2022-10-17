package routes

import (
	"github.com/gorilla/mux"
	"github.com/md/go-pro-main/internal/controllers"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/", controllers.Home).Methods("GET")                   // root
	r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")          // get all posts
	r.HandleFunc("/post/{id}", controllers.GetPostById).Methods("GET")   // get post by id
	r.HandleFunc("/post", controllers.CreatePost).Methods("POST")        // create post
	r.HandleFunc("/post/{id}", controllers.UpdatePost).Methods("PUT")    // update post by id
	r.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE") // delete post by id
}
