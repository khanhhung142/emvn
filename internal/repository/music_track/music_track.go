package musictrack_repository

import (
	"context"
	"emvn/consts"
	"emvn/database/nosql"
	"emvn/internal/model"
	"emvn/pkg/storage"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMusicTrackRepository interface {
	Create(ctx context.Context, track model.MusicTrack) (model.MusicTrack, error)
	UploadTrack(ctx context.Context, file []byte, fileName string) (string, error)
	Get(ctx context.Context, id string) (model.MusicTrack, error)
	Update(ctx context.Context, id string, track model.MusicTrack) (model.MusicTrack, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, in model.MusicTrack) ([]model.MusicTrack, error)
	GetByIDs(ctx context.Context, ids []string) ([]model.MusicTrack, error)
}

type musicTrackRepository struct {
	noSqlDB nosql.NoSQLInterface
	storage storage.StorageInterface
}

var localMusicTrackRepository IMusicTrackRepository

func InitMusicTrackRepository(noSqlDB nosql.NoSQLInterface, storage storage.StorageInterface) {
	localMusicTrackRepository = &musicTrackRepository{
		noSqlDB: noSqlDB,
		storage: storage,
	}
}

func MusicTrackRepository() IMusicTrackRepository {
	return localMusicTrackRepository
}

func (repo *musicTrackRepository) Create(ctx context.Context, track model.MusicTrack) (model.MusicTrack, error) {
	result, err := repo.noSqlDB.InsertOne(ctx, consts.MongoDBCollectionTracks, track)
	if err != nil {
		slog.Error(err.Error())
		return model.MusicTrack{}, consts.CodeInvalidRequest
	}

	return repo.Get(ctx, result.InsertedID.(primitive.ObjectID).Hex())
}

func (repo *musicTrackRepository) UploadTrack(ctx context.Context, file []byte, fileName string) (string, error) {
	path, err := repo.storage.SaveFile(file, fileName)
	if err != nil {
		slog.Error(err.Error())
		return "", consts.CodeStorageError
	}
	return path, nil
}

func (repo *musicTrackRepository) Get(ctx context.Context, id string) (model.MusicTrack, error) {
	result, err := repo.noSqlDB.FindByObjectID(ctx, consts.MongoDBCollectionTracks, id)
	if err != nil {
		slog.Error(err.Error())
		return model.MusicTrack{}, consts.CodeInternalError
	}

	var track model.MusicTrack
	err = result.Decode(&track)
	if err != nil {
		slog.Error(err.Error())
		return model.MusicTrack{}, consts.CodeMusicTrackNotFound
	}
	return track, nil
}

func (repo *musicTrackRepository) Update(ctx context.Context, id string, in model.MusicTrack) (model.MusicTrack, error) {
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "duration", Value: in.Duration},
			{Key: "link", Value: in.Link},
			{Key: "title", Value: in.Title},
			{Key: "year", Value: in.Year},
			{Key: "album", Value: in.Album},
			{Key: "artist", Value: in.Artist},
			{Key: "genre", Value: in.Genre},
		}},
	}
	_, err := repo.noSqlDB.UpdateByID(ctx, consts.MongoDBCollectionTracks, id, update)
	if err != nil {
		slog.Error(err.Error())
		return model.MusicTrack{}, consts.CodeInternalError
	}

	return repo.Get(ctx, id)
}

func (repo *musicTrackRepository) Delete(ctx context.Context, id string) error {
	music, err := repo.Get(ctx, id)
	if err != nil {
		slog.Error(err.Error())
		return consts.CodeInternalError
	}
	err = repo.noSqlDB.DeleteByID(ctx, consts.MongoDBCollectionTracks, id)
	if err != nil {
		slog.Error(err.Error())
		return consts.CodeInternalError
	}

	err = repo.storage.DeleteFile(music.Link)
	if err != nil {
		slog.Error(err.Error())
	}
	return nil
}

func (repo *musicTrackRepository) Search(ctx context.Context, in model.MusicTrack) ([]model.MusicTrack, error) {
	fiter := bson.M{}
	if in.Title != "" {
		fiter["title"] = bson.M{"$regex": in.Title, "$options": "i"}
	}
	if in.Artist != "" {
		fiter["artist"] = bson.M{"$regex": in.Artist, "$options": "i"}
	}
	if in.Album != "" {
		fiter["album"] = bson.M{"$regex": in.Album, "$options": "i"}
	}
	if in.Genre != "" {
		fiter["genre"] = bson.M{"$regex": in.Genre, "$options": "i"}
	}

	result, err := repo.noSqlDB.Find(ctx, consts.MongoDBCollectionTracks, fiter)
	if err != nil {
		slog.Error(err.Error())
		return nil, consts.CodeInternalError
	}

	var tracks []model.MusicTrack
	err = result.All(ctx, &tracks)
	if err != nil {
		slog.Error(err.Error())
		return nil, consts.CodeInternalError
	}
	return tracks, nil
}

func (repo *musicTrackRepository) GetByIDs(ctx context.Context, ids []string) ([]model.MusicTrack, error) {
	objectIDs := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			slog.Error(err.Error())
			return nil, consts.CodeInternalError
		}
		objectIDs = append(objectIDs, objectID)
	}

	fiter := bson.M{
		"_id": bson.M{"$in": objectIDs},
	}

	result, err := repo.noSqlDB.Find(ctx, consts.MongoDBCollectionTracks, fiter)
	if err != nil {
		slog.Error(err.Error())
		return nil, consts.CodeInternalError
	}

	var tracks []model.MusicTrack
	err = result.All(ctx, &tracks)
	if err != nil {
		slog.Error(err.Error())
		return nil, consts.CodeInternalError
	}
	return tracks, nil
}
