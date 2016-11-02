// comment.go
// Copyright (C) 2016 Reza Jatnika <rezajatnika@gmail.com>
//
// Distributed under terms of the MIT license.
//

package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Comment defines data structure of a Comment
type Comment struct {
	ID        bson.ObjectId `bson:"_id"`
	CreatedAt time.Time     `bson:"created_at"`
	Content   string        `bson:"content"`
}
