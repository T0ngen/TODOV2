package api

import (
	"fmt"
	"time"
	"todov2/pkg/common/db/models"

	"github.com/gin-gonic/gin"
)


const layout = "02.01.2006"


func (h *handler) GetTasks(c *gin.Context){
	fmt.Println(c.Query("search"))
	if c.Query("search") == "" {
		tasks, err := h.DB.GetAllTasks()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if tasks == nil {
			tasks = []models.Task{}
		}
		c.JSON(200, gin.H{"tasks": tasks })
		return
	}

	isTime := true
	parsedTime, err := time.Parse(layout, c.Query("search"))
	
	if err != nil {
		isTime = false
		
	}

	if isTime {
		
		parsedTimeToString := parsedTime.Format("20060102")
		tasks, err := h.DB.GetTasksByDate(parsedTimeToString)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if tasks == nil {
			tasks = []models.Task{}
		}
		c.JSON(200, gin.H{"tasks": tasks })
		return
	}
	
	if !isTime {
		fmt.Println("tut")
		tasks, err := h.DB.GetTasksLikeLtext(c.Query("search"))
		fmt.Println(tasks)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if tasks == nil {
			tasks = []models.Task{}
		}
		c.JSON(200, gin.H{"tasks": tasks })
		return
	}

	
}