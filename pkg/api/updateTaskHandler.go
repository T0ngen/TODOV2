package api

import (
	"net/http"
	"time"
	"todov2/pkg/common/db/models"
	"todov2/pkg/common/util/ndate"
	

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)



func (h *handler) UpdateTaskById(c *gin.Context){
	nowParse := time.Now().Format(expectedLayout)
	var requestBody models.Task

	{
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			logrus.WithFields(genLog(err, "ShouldBindJSON",
			"addTaskHandler")).Errorf("Unable to bind requested json to the model")
			c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to bind requested json to the model",
				})
				return
		}
	}

	

	{
		err := h.Validator.Struct(requestBody)
		fieldErrors := make(map[string]string)
		if err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {

				for _, fieldError := range validationErrors {

					fieldErrors[fieldError.Field()] = fieldError.ActualTag()
				}
			}
			logrus.WithFields(genLog(err, "AddTask",
			"addTaskHandler")).Errorf("Error validating requested input in json")
			c.JSON(http.StatusBadRequest, gin.H{
				"error":            "task title not specified, try again with title",
				})
				return
			}
	}

	
	
	
	if requestBody.Date == ""{
		requestBody.Date = time.Now().Format(expectedLayout)

	}
	if requestBody.Date != ""{
		parsedTime, err := time.Parse(expectedLayout, requestBody.Date)
	if err != nil {
		logrus.WithFields(genLog(err, "AddTask",
		"addTaskHandler")).Errorf("Error parsing date -1")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error parsing date",
		})
		return
	}
	if parsedTime.Format(expectedLayout) != requestBody.Date{
		logrus.WithFields(genLog(err, "AddTask",
		"addTaskHandler")).Errorf("Error parsing date 0")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error parsing date",
		})
		return
	}
	requestBody.Date = parsedTime.Format(expectedLayout)

	}
	now := time.Now()
	
	if requestBody.Date < nowParse{
		if requestBody.Repeat == ""{
			requestBody.Date = nowParse
		}
		if requestBody.Repeat != ""{
			dateNext, err :=  ndate.NextDate(now, requestBody.Date, requestBody.Repeat)
			if err != nil {
				logrus.WithFields(genLog(err, "AddTask",
				"addTaskHandler")).Errorf("Error parsing date 2")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error parsing date",
				})
				return
			}
			requestBody.Date = dateNext
		}
	}
	
	task, err := h.DB.GetTaskByID(requestBody.Id)
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

	err = h.DB.UpdateTaskById(requestBody)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to update task",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
	})
}