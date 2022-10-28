package router

import (
	"FP2/config"
	"FP2/controllers"
	"FP2/database"
	"FP2/middlewares"
	"FP2/repositories"
	"FP2/service"

	"github.com/gin-gonic/gin"
)

func StartAPP() *gin.Engine {
	cfg := config.LoadConfig()
	db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
	// USER
	userRepo := repositories.NewUserRepo(db)
	userService := service.NewUserService(&userRepo)
	userController := controllers.NewUserController(userService)
	// Photo
	photoRepo := repositories.NewPhotoRepo(db)
	photoService := service.NewPhotoService(&photoRepo)
	photoController := controllers.NewPhotoController(photoService)
	// Comment
	commentRepo := repositories.NewCommentRepo(db)
	commentService := service.NewCommentService(&commentRepo)
	commentController := controllers.NewCommentController(commentService)
	// Sosmed
	sosmedRepo := repositories.NewSosmedRepo(db)
	sosmedService := service.NewSosmedService(&sosmedRepo)
	sosmedController := controllers.NewSosmedController(sosmedService)

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegisterControllers)
		userRouter.POST("/login", userController.UserLoginControllers)

		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/update-user/:penggunaId", middlewares.UserAuthorization(), userController.UpdateUserController)
		userRouter.DELETE("/delete-user/:penggunaId", middlewares.UserAuthorization(), userController.DeleteUserController)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/create-photo", photoController.CreatePhotoControllers)
		photoRouter.GET("/", photoController.GetPhotoControllers)
		photoRouter.PUT("/update-photo/:photoId", middlewares.PhotoAuthorization(), photoController.UpdatePhotoControllers)
		photoRouter.DELETE("/delete-photo/:photoId", middlewares.PhotoAuthorization(), photoController.DeletePhotoControllers)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/create-comment", commentController.CreateCommentControllers)
		commentRouter.GET("/", commentController.GetCommentControllers)
		commentRouter.PUT("/update-comment/:commentId", middlewares.CommentAuthorization(), commentController.UpdateCommentControllers)
		commentRouter.DELETE("/delete-comment/:commentId", middlewares.CommentAuthorization(), commentController.DeleteCommentControllers)
	}

	sosmedRouter := router.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.POST("/create-sosmed", sosmedController.CreateSosmedControllers)
		sosmedRouter.GET("/", sosmedController.GetSosmedControllers)
		sosmedRouter.PUT("/update-sosmed/:socialMediaId", middlewares.SosmedAuthorization(), sosmedController.UpdateSosmedControllers)
		sosmedRouter.DELETE("/delete-sosmed/:socialMediaId", middlewares.SosmedAuthorization(), sosmedController.DeleteSosmedControllers)
	}

	return router
}
