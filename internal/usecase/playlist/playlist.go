package playlist_usecase

import (
	"context"
	"emvn/consts"
	"emvn/internal/model"
	musictrack_repository "emvn/internal/repository/music_track"
	playlist_repository "emvn/internal/repository/playlist"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPlaylistUsecase interface {
	Create(ctx context.Context, in model.Playlist, uid string) (PlaylistWithTracks, error)
	Get(ctx context.Context, id string) (PlaylistWithTracks, error)
	Update(ctx context.Context, id string, in model.Playlist) (PlaylistWithTracks, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, in model.Playlist) ([]model.Playlist, error)
}

type playlistUsecase struct {
	repo      playlist_repository.IPlaylistRepository
	musicRepo musictrack_repository.IMusicTrackRepository
}

// Singleton pattern
var localPlaylistUsecase IPlaylistUsecase

func InitPlaylistUsecase(repo playlist_repository.IPlaylistRepository, musicRepo musictrack_repository.IMusicTrackRepository) {
	localPlaylistUsecase = &playlistUsecase{
		repo:      repo,
		musicRepo: musicRepo,
	}
}

func PlaylistUsecase() IPlaylistUsecase {
	return localPlaylistUsecase
}

// Create a new playlist
func (usecase *playlistUsecase) Create(ctx context.Context, in model.Playlist, uid string) (PlaylistWithTracks, error) {
	in.ID = primitive.NewObjectID()
	in.CreatedBy = uid
	// checking if track_ids is valid
	// User can create a playlist without any track. they can add tracks later
	if in.TrackIDs == nil {
		in.TrackIDs = []string{}
	}

	tracks, err := usecase.musicRepo.GetByIDs(ctx, in.TrackIDs)
	if err != nil {
		return PlaylistWithTracks{}, err
	}

	// checking if all track_ids are valid
	if len(tracks) != len(in.TrackIDs) {
		return PlaylistWithTracks{}, consts.CodeMusicTrackNotFound
	}

	playlistDB, err := usecase.repo.Create(ctx, in)
	if err != nil {
		return PlaylistWithTracks{}, err
	}

	return PlaylistWithTracks{
		ID:          playlistDB.ID,
		Title:       playlistDB.Title,
		Description: playlistDB.Description,
		Genre:       playlistDB.Genre,
		CreatedBy:   playlistDB.CreatedBy,
		Tracks:      tracks,
	}, nil
}

// Get a playlist by ID
func (usecase *playlistUsecase) Get(ctx context.Context, id string) (PlaylistWithTracks, error) {
	dbPlaylist, err := usecase.repo.Get(ctx, id)
	if err != nil {
		return PlaylistWithTracks{}, err
	}

	tracks, err := usecase.musicRepo.GetByIDs(ctx, dbPlaylist.TrackIDs)
	if err != nil {
		return PlaylistWithTracks{}, err
	}

	return PlaylistWithTracks{
		ID:          dbPlaylist.ID,
		Title:       dbPlaylist.Title,
		Description: dbPlaylist.Description,
		Genre:       dbPlaylist.Genre,
		CreatedBy:   dbPlaylist.CreatedBy,
		Tracks:      tracks,
	}, nil
}

// Update a playlist by ID
func (usecase *playlistUsecase) Update(ctx context.Context, id string, in model.Playlist) (PlaylistWithTracks, error) {
	if in.TrackIDs == nil {
		in.TrackIDs = []string{}
	}

	tracks, err := usecase.musicRepo.GetByIDs(ctx, in.TrackIDs)
	if err != nil {
		return PlaylistWithTracks{}, err
	}

	// checking if all track_ids are valid
	if len(tracks) != len(in.TrackIDs) {
		return PlaylistWithTracks{}, consts.CodeMusicTrackNotFound
	}

	updatedPlaylist, err := usecase.repo.Update(ctx, id, in)
	if err != nil {
		return PlaylistWithTracks{}, err
	}

	return PlaylistWithTracks{
		ID:          updatedPlaylist.ID,
		Title:       updatedPlaylist.Title,
		Description: updatedPlaylist.Description,
		Genre:       updatedPlaylist.Genre,
		CreatedBy:   updatedPlaylist.CreatedBy,
		Tracks:      tracks,
	}, nil
}

// Delete a playlist by ID
func (usecase *playlistUsecase) Delete(ctx context.Context, id string) error {
	return usecase.repo.Delete(ctx, id)
}

// I image this function is used to search for display purposes, so we don't need to return the tracks.
// User can get tracks when they click on the playlist
func (usecase *playlistUsecase) Search(ctx context.Context, in model.Playlist) ([]model.Playlist, error) {
	playlists, err := usecase.repo.Search(ctx, in)
	if err != nil {
		return nil, err
	}
	// omitting track_ids
	for i := range playlists {
		playlists[i].TrackIDs = nil
	}

	return playlists, nil
}
