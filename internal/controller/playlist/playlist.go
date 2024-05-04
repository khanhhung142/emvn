package playlist_controller

import (
	"emvn/consts"
	"emvn/internal/model"
	playlist_usecase "emvn/internal/usecase/playlist"
	"emvn/pkg/validator"

	"github.com/gin-gonic/gin"
)

type IPlaylistController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
}

type playlistController struct {
	usecase playlist_usecase.IPlaylistUsecase
}

func NewController(usecase playlist_usecase.IPlaylistUsecase) IPlaylistController {
	return &playlistController{
		usecase: usecase,
	}
}

// CreatePlaylist swagger documentation
//
//	@Summary		Create a new playlist
//	@Description	Note that all track ids must be valid
//	@Tags			Playlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		WritePlaylistInput	true	"Playlist information"
//	@Success		200		{object}	WritePlaylistOutput
//	@Router			/playlist/create [post]
func (ctrl *playlistController) Create(c *gin.Context) {
	var in WritePlaylistInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}

	// validate track ids
	if !validateTrackIds(in.TrackIDs) {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}

	uid := c.GetString(consts.GinAuthUid)
	if uid == "" {
		c.Set(consts.GinErrorKey, consts.CodeInvalidToken)
		return
	}

	// call usecase
	newPlaylist, err := ctrl.usecase.Create(c, model.Playlist{
		Title:       in.Title,
		Description: in.Description,
		Genre:       in.Genre,
		TrackIDs:    in.TrackIDs,
	}, uid)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}

	c.Set(consts.GinResponseKey, WritePlaylistOutput{
		Title:       newPlaylist.Title,
		Description: newPlaylist.Description,
		Genre:       newPlaylist.Genre,
		CreatedBy:   newPlaylist.CreatedBy,
		Tracks:      newPlaylist.Tracks,
		ID:          newPlaylist.ID,
	})
}

// GetPlaylist swagger documentation
//
//	@Summary		Get a playlist by ID
//	@Description	Get a playlist by its ID
//	@Tags			Playlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Playlist ID"
//	@Success		200	{object}	WritePlaylistOutput
//	@Router			/playlist/get/{id} [get]
func (ctrl *playlistController) Get(c *gin.Context) {
	id := c.Param("id")
	ok := validator.IsMongoObjectId(id)
	if !ok {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}

	playlist, err := ctrl.usecase.Get(c, id)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}

	c.Set(consts.GinResponseKey, WritePlaylistOutput{
		Title:       playlist.Title,
		Description: playlist.Description,
		Genre:       playlist.Genre,
		CreatedBy:   playlist.CreatedBy,
		Tracks:      playlist.Tracks,
		ID:          playlist.ID,
	})
}

// UpdatePlaylist swagger documentation
//
//	@Summary		Update a playlist by ID
//	@Description	Update a playlist by its ID
//	@Tags			Playlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		string				true	"Playlist ID"
//	@Param			request	body		WritePlaylistInput	true	"Playlist information"
//	@Success		200		{object}	WritePlaylistOutput
//	@Router			/playlist/update/{id} [put]
func (ctrl *playlistController) Update(c *gin.Context) {
	var in WritePlaylistInput
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

	// validate track ids
	if !validateTrackIds(in.TrackIDs) {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}

	// call usecase
	newPlaylist, err := ctrl.usecase.Update(c, id, model.Playlist{
		Title:       in.Title,
		Description: in.Description,
		Genre:       in.Genre,
		TrackIDs:    in.TrackIDs,
	})
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}

	c.Set(consts.GinResponseKey, WritePlaylistOutput{
		Title:       newPlaylist.Title,
		Description: newPlaylist.Description,
		Genre:       newPlaylist.Genre,
		CreatedBy:   newPlaylist.CreatedBy,
		Tracks:      newPlaylist.Tracks,
		ID:          newPlaylist.ID,
	})
}

// DeletePlaylist swagger documentation
//
//	@Summary		Delete a playlist by ID
//	@Description	Delete a playlist by its ID
//	@Tags			Playlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Playlist ID"
//	@Success		200	{object}	TempOut
//	@Router			/playlist/delete/{id} [delete]
func (ctrl *playlistController) Delete(c *gin.Context) {
	id := c.Param("id")
	ok := validator.IsMongoObjectId(id)
	if !ok {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		return
	}

	err := ctrl.usecase.Delete(c, id)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}

	c.Set(consts.GinResponseKey, gin.H{"success": true})
}

// SearchPlaylist swagger documentation
//
//	@Summary		Search playlists based on criteria
//	@Description	Search playlists based on title, description, and genre
//	@Tags			Playlist
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			title		query		string	false	"Playlist title"
//	@Param			description	query		string	false	"Playlist description"
//	@Param			genre		query		string	false	"Playlist genre"
//	@Success		200			{object}	[]model.Playlist
//	@Router			/playlist/search [get]
func (ctrl *playlistController) Search(c *gin.Context) {
	var in SearchPlaylistInput
	if err := c.ShouldBindQuery(&in); err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}

	// call usecase
	playlists, err := ctrl.usecase.Search(c, model.Playlist{
		Title:       in.Title,
		Description: in.Description,
		Genre:       in.Genre,
	})
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}

	c.Set(consts.GinResponseKey, playlists)
}

func validateTrackIds(trackIds []string) bool {
	for _, id := range trackIds {
		if !validator.IsMongoObjectId(id) {
			return false
		}
	}
	return true
}
