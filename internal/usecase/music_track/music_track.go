package musictrack_usecase

import (
	"context"
	"emvn/consts"
	"emvn/internal/model"
	musictrack_repository "emvn/internal/repository/music_track"
	"log/slog"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMusicTrackUsecase interface {
	CreateMusicTrack(ctx context.Context, in model.MusicTrack) (model.MusicTrack, error)
	UploadTrack(ctx context.Context, file *multipart.FileHeader) (string, error)
	GetMusicTrack(ctx context.Context, id string) (model.MusicTrack, error)
	UpdateMusicTrack(ctx context.Context, id string, in model.MusicTrack) (model.MusicTrack, error)
	DeleteMusicTrack(ctx context.Context, id string) error
	SearchMusicTrack(ctx context.Context, in model.MusicTrack) ([]model.MusicTrack, error)
}

type musicTrackUsecase struct {
	musicTrackRepo musictrack_repository.IMusicTrackRepository
}

var localMusicTrackUsecase IMusicTrackUsecase

func InitMusicTrackUsecase(musicTrackRepo musictrack_repository.IMusicTrackRepository) {
	localMusicTrackUsecase = &musicTrackUsecase{
		musicTrackRepo: musicTrackRepo,
	}
}

func MusicTrackUsecase() IMusicTrackUsecase {
	return localMusicTrackUsecase
}

func (uc *musicTrackUsecase) CreateMusicTrack(ctx context.Context, in model.MusicTrack) (model.MusicTrack, error) {
	in.ID = primitive.NewObjectID()
	return uc.musicTrackRepo.Create(ctx, in)
}

func (uc *musicTrackUsecase) UploadTrack(ctx context.Context, file *multipart.FileHeader) (string, error) {
	fileOpen, err := file.Open()
	if err != nil {
		slog.Error(err.Error())
		return "", consts.CodeFileInvalid
	}

	fileBytes := make([]byte, file.Size)
	_, err = fileOpen.Read(fileBytes)
	if err != nil {
		slog.Error(err.Error())
		return "", consts.CodeFileInvalid
	}

	return uc.musicTrackRepo.UploadTrack(ctx, fileBytes, file.Filename)
}

func (uc *musicTrackUsecase) GetMusicTrack(ctx context.Context, id string) (model.MusicTrack, error) {
	return uc.musicTrackRepo.Get(ctx, id)
}

func (uc *musicTrackUsecase) UpdateMusicTrack(ctx context.Context, id string, in model.MusicTrack) (model.MusicTrack, error) {
	return uc.musicTrackRepo.Update(ctx, id, in)
}

func (uc *musicTrackUsecase) DeleteMusicTrack(ctx context.Context, id string) error {
	return uc.musicTrackRepo.Delete(ctx, id)
}

func (uc *musicTrackUsecase) SearchMusicTrack(ctx context.Context, in model.MusicTrack) ([]model.MusicTrack, error) {
	return uc.musicTrackRepo.Search(ctx, in)
}
