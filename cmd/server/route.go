package server

import (
	auth_controller "emvn/internal/controller/auth"
	musictrack_controller "emvn/internal/controller/music_track"
	playlist_controller "emvn/internal/controller/playlist"
	auth_usecase "emvn/internal/usecase/auth"
	musictrack_usecase "emvn/internal/usecase/music_track"
	playlist_usecase "emvn/internal/usecase/playlist"
	"emvn/middlewares"

	doc "emvn/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitHandler() *gin.Engine {
	// Init gin
	r := gin.Default()

	// Add middlewares
	r.Use(middlewares.LogMiddleware())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ResponseMiddleware())

	// Add routes

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	authController := auth_controller.NewController(auth_usecase.AuthUsecase())

	authGroup := r.Group("/auth")
	authGroup.POST("/signup", authController.SignUp)
	authGroup.POST("/signin", authController.SignIn)

	mucisTrackController := musictrack_controller.NewController(musictrack_usecase.MusicTrackUsecase())

	musicTrackGroup := r.Group("/music_track", middlewares.AuthMiddleware())
	musicTrackGroup.POST("/create", mucisTrackController.Create)
	musicTrackGroup.POST("/upload", mucisTrackController.UploadTrack)
	musicTrackGroup.GET("/get/:id", mucisTrackController.Get)
	musicTrackGroup.PUT("/update/:id", mucisTrackController.Update)
	musicTrackGroup.DELETE("/delete/:id", mucisTrackController.Delete)
	musicTrackGroup.GET("/search", mucisTrackController.Search)

	playlistController := playlist_controller.NewController(playlist_usecase.PlaylistUsecase())

	playlistGroup := r.Group("/playlist", middlewares.AuthMiddleware())
	playlistGroup.POST("/create", playlistController.Create)
	playlistGroup.GET("/get/:id", playlistController.Get)
	playlistGroup.PUT("/update/:id", playlistController.Update)
	playlistGroup.DELETE("/delete/:id", playlistController.Delete)
	playlistGroup.GET("/search", playlistController.Search)

	// Swagger
	doc.SwaggerInfo.Title = "EMVN API"
	doc.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
