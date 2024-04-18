package api

import (
	"net/http"
	"time"
	
	"todov2/pkg/common/db/models"
	"todov2/pkg/common/util/ndate"

	"github.com/gin-gonic/gin"
)


func (h *handler) TaskDone (c *gin.Context){

	id := c.Query("id")

	if id == ""{
		c.JSON(400, gin.H{"error": "id is required"})
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

	if task.Repeat == ""{
		err  := h.DB.DeleteTaskById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to delete task",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	newDate, err := ndate.NextDate(time.Now(), task.Date,task.Repeat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete task",
		})
		return
	}

	newTask := models.Task{
		Id: id,
		Date: newDate,
		Title: task.Title,
		Comment: task.Comment,
		Repeat: task.Repeat,
	}
	err = h.DB.UpdateTaskById(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}