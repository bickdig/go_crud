// post.go
// Copyright (C) 2016 Reza Jatnika <rezajatnika@gmail.com>
//
// Distributed under terms of the MIT license.
//

package models

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Post define the data structure of a Post
type Post struct {
	ID        bson.ObjectId `bson:"_id"`
	CreatedAt time.Time     `bson:"created_at"`
	Title     string        `bson:"title"`
	Content   string        `bson:"content"`
	Comments  []comment     `bson:"comments"`
}

type comment struct {
	ID        bson.ObjectId `bson:"_id"`
	CreatedAt time.Time     `bson:"created_at"`
	Content   string        `bson:"content"`
}

// PostCollection define the data structure of a Post
type PostCollection struct {
	collection *mgo.Collection
}

func getCollection() *mgo.Collection {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return session.DB("go_crud").C("posts")
}

// NewPostCollection returns PostCollection struct
func NewPostCollection() *PostCollection {
	return &PostCollection{getCollection()}
}

// All returns all posts
func (pc *PostCollection) All() []Post {
	posts := []Post{}
	pc.collection.Find(nil).All(&posts)
	return posts
}

// FindID returns a post by its object id
func (pc *PostCollection) FindID(oid bson.ObjectId) *Post {
	post := &Post{}
	pc.collection.FindId(oid).One(&post)
	return post
}

// Insert inserts a post to database and returns itself
func (pc *PostCollection) Insert(post *Post) *Post {
	pc.collection.Insert(&post)
	return post
}

// RemoveID removes a post by its object id
func (pc *PostCollection) RemoveID(oid bson.ObjectId) {
	pc.collection.RemoveId(oid)
}

// UpdateID updates a post by its object id
func (pc *PostCollection) UpdateID(oid bson.ObjectId, post *Post) {
	pc.collection.UpdateId(oid, &post)
}

// Path returns post path
func (p *Post) Path() string {
	return "/posts/" + p.ID.Hex()
}
