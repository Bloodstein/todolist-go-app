package controllers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", handler.createList)
			lists.GET("/", handler.getAllLists)
			lists.GET("/:id", handler.getListById)
			lists.PUT("/:id", handler.updateList)
			lists.DELETE("/:id", handler.deleteList)

			items := lists.Group("/items")
			{
				items.POST("/", handler.createItem)
				items.GET("/", handler.getAllItems)
				items.GET("/:item_id", handler.getItemById)
				items.PUT("/:item_id", handler.updateItem)
				items.DELETE("/:item_id", handler.deleteItem)
			}
		}
	}

	return router
}
