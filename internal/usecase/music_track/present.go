package musictrack_usecase

type CreateMusicTrackInput struct {
	Title    string   `bson:"title" json:"title"`
	Artists  []string `bson:"artists" json:"artists"`
	Album    string   `bson:"album" json:"album"`
	Genres   []string `bson:"genres" json:"genres"`
	Year     int      `bson:"year" json:"year"`
	Duration int      `bson:"duration" json:"duration"`
	Link     string   `bson:"link" json:"link"`
}
