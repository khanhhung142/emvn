package playlist_repository

import (
	"context"
	"emvn/consts"
	"emvn/database/nosql"
	"emvn/internal/model"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPlaylistRepository interface {
	Create(ctx context.Context, playlist model.Playlist) (model.Playlist, error)
	Get(ctx context.Context, id string) (model.Playlist, error)
	Update(ctx context.Context, id string, playlist model.Playlist) (model.Playlist, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, in model.Playlist) ([]model.Playlist, error)
}

type playlistRepository struct {
	noSqlDB nosql.NoSQLInterface
}

// Singleton pattern
var localPlaylistRepository IPlaylistRepository

func InitPlaylistRepository(noSqlDB nosql.NoSQLInterface) {
	localPlaylistRepository = &playlistRepository{
		noSqlDB: noSqlDB,
	}
}

func PlaylistRepository() IPlaylistRepository {
	return localPlaylistRepository
}

// Create a new playlist
func (repo *playlistRepository) Create(ctx context.Context, playlist model.Playlist) (model.Playlist, error) {
	result, err := repo.noSqlDB.InsertOne(ctx, consts.MongoDBCollectionPlaylists, playlist)
	if err != nil {
		slog.Error(err.Error())
		return model.Playlist{}, consts.CodeInternalError
	}

	return repo.Get(ctx, result.InsertedID.(primitive.ObjectID).Hex())
}

// Get a playlist by ID
func (repo *playlistRepository) Get(ctx context.Context, id string) (model.Playlist, error) {
	result, err := repo.noSqlDB.FindByObjectID(ctx, consts.MongoDBCollectionPlaylists, id)
	if err != nil {
		slog.Error(err.Error())
		return model.Playlist{}, consts.CodeInternalError
	}

	var playlist model.Playlist
	err = result.Decode(&playlist)
	if err != nil {
		slog.Error(err.Error())
		return model.Playlist{}, consts.CodeInternalError
	}

	return playlist, nil
}

// Update a playlist by ID
func (repo *playlistRepository) Update(ctx context.Context, id string, playlist model.Playlist) (model.Playlist, error) {
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: playlist.Title},
			{Key: "genre", Value: playlist.Genre},
			{Key: "description", Value: playlist.Description},
			{Key: "track_ids", Value: playlist.TrackIDs},
		}},
	}

	_, err := repo.noSqlDB.UpdateByID(ctx, consts.MongoDBCollectionPlaylists, id, update)
	if err != nil {
		slog.Error(err.Error())
		return model.Playlist{}, consts.CodeInternalError
	}

	return repo.Get(ctx, id)
}

// Delete a playlist by ID
func (repo *playlistRepository) Delete(ctx context.Context, id string) error {
	return repo.noSqlDB.DeleteByID(ctx, consts.MongoDBCollectionPlaylists, id)
}

// Search playlists
func (repo *playlistRepository) Search(ctx context.Context, in model.Playlist) ([]model.Playlist, error) {
	fiter := bson.M{}
	if in.Title != "" {
		fiter["title"] = bson.M{"$regex": in.Title, "$options": "i"}
	}
	if in.Description != "" {
		fiter["description"] = bson.M{"$regex": in.Description, "$options": "i"}
	}
	if in.Genre != "" {
		fiter["genre"] = bson.M{"$regex": in.Genre, "$options": "i"}
	}

	cursor, err := repo.noSqlDB.Find(ctx, consts.MongoDBCollectionPlaylists, fiter)
	if err != nil {
		slog.Error(err.Error())
		return nil, consts.CodeInternalError
	}

	var playlists []model.Playlist
	err = cursor.All(ctx, &playlists)
	if err != nil {
		slog.Error(err.Error())
		return nil, consts.CodeInternalError
	}

	return playlists, nil
}
