package api

import (
	"net/http"
	"todov2/pkg/common/db/models"

	"github.com/gin-gonic/gin"
)




func (h *handler) DeleteTaskById(c *gin.Context){
	id := c.Query("id")
	if id == ""{
		c.JSON(400, gin.H{"message": "id is empty"})
		return
	}
	task, err := h.DB.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get task",
		})
		return
	}
	nilTask := models.Task{}
	if task == nilTask{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get task",
		})
	}

	err =h.DB.DeleteTaskById(id)
	if err != nil{
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}


	c.JSON(200, gin.H{})
}