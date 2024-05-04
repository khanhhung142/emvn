package playlist_controller

import (
	"emvn/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WritePlaylistInput struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	TrackIDs    []string `json:"track_ids"`
	Genre       string   `json:"genre" binding:"required"`
}

type WritePlaylistOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Genre       string             `json:"genre"`
	CreatedBy   string             `json:"created_by"`
	Tracks      []model.MusicTrack `json:"tracks"`
}

type SearchPlaylistInput struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Genre       string `form:"genre"`
}

type TempOut struct {
	Success bool `json:"success"`
}
