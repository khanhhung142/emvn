package consts

type NoSQLCollection string

const (
	MongoDBCollectionUsers     NoSQLCollection = "users"
	MongoDBCollectionTracks    NoSQLCollection = "tracks"
	MongoDBCollectionPlaylists NoSQLCollection = "playlists"
)

func (m NoSQLCollection) String() string {
	return string(m)
}
