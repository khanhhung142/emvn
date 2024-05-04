package server

import (
	"emvn/database/nosql/mongodb"
	musictrack_repository "emvn/internal/repository/music_track"
	playlist_repository "emvn/internal/repository/playlist"
	user_repository "emvn/internal/repository/user"
	auth_usecase "emvn/internal/usecase/auth"
	musictrack_usecase "emvn/internal/usecase/music_track"
	playlist_usecase "emvn/internal/usecase/playlist"
	"emvn/pkg/storage/local"
)

func Register() {
	noSqlDB := mongodb.MongoDBClient()
	localStorage := local.Storage()

	user_repository.InitUserRepository(noSqlDB)
	auth_usecase.InitAuthUsecase(user_repository.UserRepository())

	musictrack_repository.InitMusicTrackRepository(noSqlDB, localStorage)
	musictrack_usecase.InitMusicTrackUsecase(musictrack_repository.MusicTrackRepository())

	playlist_repository.InitPlaylistRepository(noSqlDB)
	playlist_usecase.InitPlaylistUsecase(playlist_repository.PlaylistRepository(), musictrack_repository.MusicTrackRepository())
}
