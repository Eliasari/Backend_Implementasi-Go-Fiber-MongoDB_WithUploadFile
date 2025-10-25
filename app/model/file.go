package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Filename    string             `bson:"filename" json:"filename"`
	Path        string             `bson:"path" json:"path"`
	ContentType string             `bson:"content_type" json:"content_type"`
	Size        int64              `bson:"size" json:"size"`
	Type        string             `bson:"type" json:"type"` // photo | certificate
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

type FileWithUser struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Filename    string             `bson:"filename" json:"filename"`
	Path        string             `bson:"path" json:"path"`
	ContentType string             `bson:"content_type" json:"content_type"`
	Size        int64              `bson:"size" json:"size"`
	Type        string             `bson:"type" json:"type"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	User        []User             `bson:"user" json:"user"`
}
