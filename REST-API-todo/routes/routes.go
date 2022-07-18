package routes

import (
	"github.com/alexparco/rest-api-todo/config"
	controller "github.com/alexparco/rest-api-todo/controllers"
	"github.com/alexparco/rest-api-todo/database"
	"github.com/alexparco/rest-api-todo/middlewares"
	"github.com/alexparco/rest-api-todo/model"
	"github.com/gin-gonic/gin"
)

func NewRouter(config *config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// new client
	db := database.NewDbClient(config.Mysql)

	tr := model.NewTaskRepo(db)
	tc := controller.NewTaskController(tr)

	// set cors
	router.Use(middlewares.CORS())

	v1 := router.Group("/v1/api")
	{
		v1.GET("/tasks", tc.GetTasks)
		v1.GET("/tasks/:id", tc.GetTask)
		v1.POST("/tasks", tc.CreateTask)
		v1.PUT("/tasks/:id", tc.UpdateTask)
		v1.DELETE("/tasks/:id", tc.DeleteTask)
	}

	return router
}
