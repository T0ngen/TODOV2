package api

import (
	"net/http"
	"todov2/pkg/api/mappers"
	hashed "todov2/pkg/common/util/hash"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func (h *handler)SignIn (c *gin.Context){
	var requestBody mappers.SignUpForm

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
	if h.config.Password != requestBody.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неверный пароль",
		})
		return
	}
	token := hashed.HashPassword(h.config.Password)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	
	

	

}