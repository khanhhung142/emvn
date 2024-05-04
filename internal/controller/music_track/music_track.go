package musictrack_controller

import (
	"emvn/consts"
	"emvn/internal/model"
	musictrack_usecase "emvn/internal/usecase/music_track"
	"emvn/pkg/validator"
	"log"

	"github.com/gin-gonic/gin"
)

type IMusicTrackController interface {
	Create(c *gin.Context)
	UploadTrack(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
}

type musicTrackController struct {
	musicTrackUsecase musictrack_usecase.IMusicTrackUsecase
}

func NewController(musicTrackUsecase musictrack_usecase.IMusicTrackUsecase) IMusicTrackController {
	return &musicTrackController{
		musicTrackUsecase: musicTrackUsecase,
	}
}

// CreateMusicTrack swagger documentation
//
//	@Summary		Create a new music track
//	@Description	Create a new music track with the given information. Must upload the track file separately
//	@Tags			Music Track
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		WriteMusicTrackInput	true	"Music track information"
//	@Success		200		{object}	WriteMusicTrackOutput
//	@Router			/music_track/create [post]
func (ctrl *musicTrackController) Create(c *gin.Context) {
	// validate request
	var in WriteMusicTrackInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}

	// call usecase
	newTrack, err := ctrl.musicTrackUsecase.CreateMusicTrack(c, model.MusicTrack{
		Artist:   in.Artist,
		Album:    in.Album,
		Genre:    in.Genre,
		Year:     in.Year,
		Title:    in.Title,
		Duration: in.Duration,
		Link:     in.Link,
	})
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	c.Set(consts.GinResponseKey, WriteMusicTrackOutput{
		MusicTrack: newTrack,
	})
}

// UploadTrack swagger documentation
//	@Summary		Upload a music track
//	@Description	Upload a music track file
//	@Tags			Music Track
//	@Produce		json
//	@Security		BearerAuth
//	@Param			file	formData	file	true	"Music track file"
//	@Success		200		{object}	UploadTrackOutput
//	@Router			/music_track/upload [post]
func (ctrl *musicTrackController) UploadTrack(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Set(consts.GinErrorKey, consts.CodeFileNotFound)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}
	filePath, err := ctrl.musicTrackUsecase.UploadTrack(c, file)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	c.Set(consts.GinResponseKey, UploadTrackOutput{
		FilePath: filePath,
	})
}

// GetMusicTrack swagger documentation
//	@Summary		Get a music track by ID
//	@Description	Get a music track by its ID
//	@Tags			Music Track
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Music track ID"
//	@Success		200	{object}	model.MusicTrack
//	@Router			/music_track/get/{id} [get]
func (ctrl *musicTrackController) Get(c *gin.Context) {
	id := c.Param("id")
	ok := validator.IsMongoObjectId(id)
	if !ok {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}
	track, err := ctrl.musicTrackUsecase.GetMusicTrack(c, id)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	c.Set(consts.GinResponseKey, track)
}

// UpdateMusicTrack swagger documentation
//	@Summary		Update a music track
//	@Description	Update a music track with the given information
//	@Tags			Music Track
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		string					true	"Music track ID"
//	@Param			request	body		WriteMusicTrackInput	true	"Music track information"
//	@Success		200		{object}	WriteMusicTrackOutput
//	@Router			/music_track/update/{id} [put]
func (ctrl *musicTrackController) Update(c *gin.Context) {
	// validate request
	var in WriteMusicTrackInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}

	id := c.Param("id")
	ok := validator.IsMongoObjectId(id)
	if !ok {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}

	// call usecase
	newTrack, err := ctrl.musicTrackUsecase.UpdateMusicTrack(c, id, model.MusicTrack{
		Artist:   in.Artist,
		Album:    in.Album,
		Genre:    in.Genre,
		Year:     in.Year,
		Title:    in.Title,
		Duration: in.Duration,
		Link:     in.Link,
	})
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	c.Set(consts.GinResponseKey, WriteMusicTrackOutput{
		MusicTrack: newTrack,
	})
}

// DeleteMusicTrack swagger documentation
//	@Summary		Delete a music track
//	@Description	Delete a music track by its ID
//	@Tags			Music Track
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Music track ID"
//	@Success		200	{object}	TempOut
//	@Router			/music_track/delete/{id} [delete]
func (ctrl *musicTrackController) Delete(c *gin.Context) {
	id := c.Param("id")
	ok := validator.IsMongoObjectId(id)
	if !ok {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}
	err := ctrl.musicTrackUsecase.DeleteMusicTrack(c, id)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	c.Set(consts.GinResponseKey, gin.H{
		"success": true,
	})
}

// SearchMusicTrack swagger documentation
//	@Summary		Search music tracks
//	@Description	Search music tracks based on the provided criteria
//	@Tags			Music Track
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			artist	query		string	false	"Artist name"
//	@Param			album	query		string	false	"Album name"
//	@Param			genre	query		string	false	"Genre"
//	@Param			title	query		string	false	"Title"
//	@Success		200		{object}	[]model.MusicTrack
//	@Router			/music_track/search [get]
func (ctrl *musicTrackController) Search(c *gin.Context) {
	// validate request
	var in SearchMusicTrackInput
	err := c.ShouldBindQuery(&in)
	if err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}
	log.Println(in)
	// call usecase
	tracks, err := ctrl.musicTrackUsecase.SearchMusicTrack(c, model.MusicTrack{
		Artist: in.Artist,
		Album:  in.Album,
		Genre:  in.Genre,
		Title:  in.Title,
	})
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	c.Set(consts.GinResponseKey, tracks)
}
