package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/researchlab/gbp/http-server/middleware"
)

var di = &DomainInfo{
	Id:             time.Now().Unix(),
	Name:           "mike.json",
	CustId:         801021,
	SrcIp:          "127.0.0.1",
	LogFmt:         2,
	LogInterval:    5,
	LogWild:        1,
	Type:           2,
	HType:          1,
	LogLevel:       3,
	BitRate:        10,
	CostWithParent: 12,
}

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	d := r.Group("/domain")
	{
		d.POST("/update", Update)
		d.GET("/read", Read)
	}
	s := r.Group("/api/domain")
	{
		s.GET("/read", StockRead)
	}
	r.Run(":8000")
}

func StockRead(c *gin.Context) {

	// c.Query() 等同于 c.Request.URL.Query().Get()
	//name := c.Query("name")
	//name2 := c.Request.URL.Query().Get("name")

	// 设置默认参数,如果job参数不存在,默认为程序员
	code := c.DefaultQuery("code", "-1")
	fmt.Println("stock code:", code)
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  "10,11,12,13,14,15,16,17",
	})

}

func Update(c *gin.Context) {
	timestamp := c.DefaultPostForm("timestamp", "0")
	param := c.DefaultPostForm("param", "")
	token := c.DefaultPostForm("token", "")
	fmt.Println("timestamp:", timestamp, " token:", token)
	_di := DomainInfo{}
	err := json.Unmarshal([]byte(param), &_di)
	if err != nil {
		ResponseError(c, err)
		return
	}
	di.Update(_di)
	ResponseSuccess(c, di)
}

func Read(c *gin.Context) {
	ResponseSuccess(c, di)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"error": nil,
	})
}

func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
		"data":  nil,
	})
}
