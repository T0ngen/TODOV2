package api

import (
	"net/http"
	"todov2/pkg/api/mappers"
	"todov2/pkg/common/db"
	"todov2/pkg/common/db/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DBImpl interface{
	AddTaskToDb(task mappers.NewTask) (int, error)
	GetAllTasks() ([]models.Task, error)
	GetTasksByDate(date string) ([]models.Task, error)
	GetTasksLikeLtext(text string) ([]models.Task, error)
	GetTaskByID(id string) (models.Task, error)
	UpdateTaskById(task models.Task) error
	DeleteTaskById(taskId string) error
}

type handler struct {
	Validator *validator.Validate
	DB DBImpl
}


const (
	webDir string = "./web"
)

func RegisterRouter(r *gin.Engine, validator *validator.Validate, db *db.Db) {
    h := &handler{Validator: validator, DB: db}
	
    routes := r.Group("/api")
	
    routes.GET("/nextdate", h.NextDate)
    routes.POST("/task", h.AddTask)
	routes.GET("/tasks", h.GetTasks)
	routes.GET("/task/",  h.GetTaskById)
	routes.PUT("/task/", h.UpdateTaskById)
	routes.DELETE("/task/", h.DeleteTaskById)
	routes.POST("/task/done/", h.TaskDone)
	
	r.NoRoute(func(c *gin.Context) {
		http.FileServer(http.Dir(webDir)).ServeHTTP(c.Writer, c.Request)
	})
   
	
   

}


