package main

import (
	"github.com/CoderSamYhc/gin-web/db"
	"github.com/CoderSamYhc/gin-web/routers"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	db.InitDB()
	defer db.DB.Close()
	// 路由
	routers.Init(r)
	r.Run()
}
