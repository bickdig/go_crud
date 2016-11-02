// post_controller.go
// Copyright (C) 2016 Reza Jatnika <rezajatnika@gmail.com>
//
// Distributed under terms of the MIT license.
//

package controllers

import (
	"net/http"
	"time"

	"github.com/bickdig/go_crud/app/models"
	"github.com/bickdig/go_crud/app/views"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
)

// PostController holds the PostController methods
type PostController struct {
	session *sessions.Session
}

// Global varialbes for this package
var (
	postCollection = models.NewPostCollection()
	storeSession   = sessions.NewCookieStore([]byte("something-dark-arts"))
)

func getSession(r *http.Request, name string) *sessions.Session {
	session, err := storeSession.Get(r, name)
	if err != nil {
		panic(err)
	}
	return session
}

// NewPostController returns a new PostController struct
func NewPostController() *PostController {
	return &PostController{getSession()}
}

// New renders new template
func (uc PostController) New(w http.ResponseWriter, r *http.Request) {
	views.Render(w, "post/new", nil)

}

// Show renders show template
func (uc PostController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Check if id is an ObjectIdHex
	if !bson.IsObjectIdHex(vars["id"]) {
		w.Write([]byte("404 page not found"))
		return
	}

	post := postCollection.FindID(bson.ObjectIdHex(vars["id"]))
	views.Render(w, "post/show", post)
}

// Index renders index template
func (uc PostController) Index(w http.ResponseWriter, r *http.Request) {
	posts := postCollection.All()
	views.Render(w, "post/index", &posts)
}

// Create inserts new post and redirect to its path
func (uc PostController) Create(w http.ResponseWriter, r *http.Request) {
	// Parsing form values
	r.ParseForm()

	// Build a new post from form values
	post := &models.Post{
		ID:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
	}

	// Form validations
	if len(r.FormValue("title")) == 0 || len(r.FormValue("content")) == 0 {
		views.Render(w, "post/new", post)
		return
	}

	// Insert post to collection
	postCollection.Insert(post)

	// Redirect to post URL
	http.Redirect(w, r, post.Path(), 301)
}

// Edit renders edit template
func (uc *PostController) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Check if id is an ObjectIdHex
	if !bson.IsObjectIdHex(vars["id"]) {
		w.Write([]byte("404 page not found"))
		return
	}

	post := postCollection.FindID(bson.ObjectIdHex(vars["id"]))
	views.Render(w, "post/edit", post)
}

// Update inserts new post and redirect to its path
func (uc PostController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post := postCollection.FindID(bson.ObjectIdHex(vars["id"]))

	// Parsing form values
	r.ParseForm()

	// Build a new post from form values
	post.Title = r.FormValue("title")
	post.Content = r.FormValue("content")

	// Form validations
	if len(r.FormValue("title")) == 0 || len(r.FormValue("content")) == 0 {
		views.Render(w, "post/edit", post)
		return
	}

	// Insert post to collection
	postCollection.UpdateID(post.ID, post)

	// Redirect to post URL
	http.Redirect(w, r, post.Path(), 301)
}
