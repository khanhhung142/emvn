package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Because time limit and scope of the project, we will directly store the link of the music track
// But in my idea, we should create a cdn service
// The link field bellow will store object id of the file in the storage
// Then in cdn service we have a database to map object id to object key in the storage
// And return pre-signed url to the client
type MusicTrack struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Title    string             `bson:"title" json:"title"`
	Artist   string             `bson:"artist" json:"artist"`
	Album    string             `bson:"album" json:"album"`
	Genre    string             `bson:"genre" json:"genre"`
	Year     int                `bson:"year" json:"year"`
	Duration int                `bson:"duration" json:"duration"`
	Link     string             `bson:"link" json:"link"` // URL or local file path, get from storage
}
