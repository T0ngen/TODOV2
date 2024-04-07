package main

import (

	"todov2/pkg/api"
	"todov2/pkg/common/config"
	"todov2/pkg/common/db"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

func main() {
	var conf config.Config

    router := gin.Default()
    err := conf.InitConfig()
	if err != nil {
		logrus.Fatal("Can't set up configuration file")
	}
	validate := validator.New()
	db, err := db.NewDb(conf.DbPath)
	if err != nil {
		logrus.Fatal("Can't set up database")
	}
	



	router.GET("/", FileServer)
	
    api.RegisterRouter(router, validate, db)
	

	
	
	router.Run(":" + conf.Port)


	
}

const webPath = "./web/"

func FileServer(c *gin.Context) {
	c.File(webPath + c.Request.URL.Path)
} 

