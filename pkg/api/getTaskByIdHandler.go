package api

import (
	

	"github.com/gin-gonic/gin"
)



func (h *handler) GetTaskById(c *gin.Context){
	id := c.Query("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "ID not specified"})
		return
	}
	task, err := h.DB.GetTaskByID(id)
	if err != nil {
		c.JSON(400, gin.H{"error":"Task not found"})
		return
	}

	c.JSON(200, task)
}