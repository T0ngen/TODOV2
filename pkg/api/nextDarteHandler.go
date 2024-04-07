package api

import (
	"fmt"
	"time"
	"todov2/pkg/common/util/ndate"

	"github.com/gin-gonic/gin"
)




func (h *handler) NextDate(c *gin.Context){
	nowQuery := c.Query("now")
        query := c.Query("date")
        repeat := c.Query("repeat")
		fmt.Println(nowQuery)

		dateFormat, err := time.Parse("20060102", nowQuery)
		if err != nil {
			c.Writer.Write([]byte("error"))
			return
		}
		fmt.Println(dateFormat)

        date, err:= ndate.NextDate(dateFormat, query, repeat)
		if err != nil {
			c.Writer.Write([]byte("error"))
		}
        c.Writer.Write([]byte(date))
}