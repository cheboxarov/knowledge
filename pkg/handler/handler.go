package handler

import (
	"github.com/cheboxarov/knowledge/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		articles := api.Group("/articles")
		{
			articles.GET("/", h.getAllArticles)
			articles.GET("/:id", h.getArticle)
			articles.POST("/", h.createArticle)
			articles.DELETE("/:id", h.deleteArticle)
			articles.PATCH("/:id", h.updateArticle)
		}

		articlesGroups := api.Group("/articles-groups")
		{
			articlesGroups.GET("/", h.getAllArticlesGroups)
			articlesGroups.GET("/:id", h.getArticlesGroup)
			articlesGroups.POST("/", h.createArticlesGroup)
			articlesGroups.DELETE("/:id", h.deleteArticlesGroup)
			articlesGroups.PATCH("/:id", h.updateArticlesGroup)
		}
	}
	return router
}
