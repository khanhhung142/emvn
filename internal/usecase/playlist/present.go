package playlist_usecase

import (
	"emvn/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlaylistWithTracks struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Genre       string             `json:"genre"`
	Tracks      []model.MusicTrack `json:"tracks"`
	CreatedBy   string             `json:"created_by"` // uid of the user who created the playlist
}
