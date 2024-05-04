package musictrack_controller

import "emvn/internal/model"

type WriteMusicTrackInput struct {
	Title    string `json:"title" binding:"required"`
	Artist   string `json:"artist" binding:"required"`
	Album    string `json:"album" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
	Year     int    `json:"year" binding:"required,min=1"`
	Duration int    `json:"duration" binding:"required,min=1"`
	Link     string `json:"link" binding:"required"`
}

type WriteMusicTrackOutput struct {
	model.MusicTrack
}

type UploadTrackOutput struct {
	FilePath string `json:"file_path"`
}

type SearchMusicTrackInput struct {
	Title  string `form:"title"`
	Artist string `form:"artist"`
	Album  string `form:"album"`
	Genre  string `form:"genre"`
}

type TempOut struct {
	Success bool `json:"success"`
}
