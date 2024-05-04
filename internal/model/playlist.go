package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Genre       string             `bson:"genre" json:"genre"`
	TrackIDs    []string           `bson:"track_ids" json:"track_ids,omitempty"`
	CreatedBy   string             `bson:"created_by" json:"created_by"` // uid of the user who created the playlist
}
