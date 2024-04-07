package api

import (
	
	"todov2/pkg/api/mappers"
	"todov2/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type DBImpl interface{
	AddTaskToDb(task mappers.NewTask) (int, error)
}

type handler struct {
	Validator *validator.Validate
	DB DBImpl
}




func RegisterRouter(r *gin.Engine,validator *validator.Validate,db *db.Db) {

	h := &handler{Validator: validator, DB: db}
	
	routes := r.Group("/api/")
	routes.GET("/nextdate/",h.NextDate)
	routes.POST("/task", h.AddTask)
}
