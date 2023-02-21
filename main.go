package main

import (
	"github.com/Limeng-svg/ByteDance/dal"
	"github.com/Limeng-svg/ByteDance/dal/query"
	"github.com/Limeng-svg/ByteDance/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := dal.InitDB(); err != nil {
		os.Exit(-1)
	}
	query.SetDefault(dal.DB)
	r := gin.Default()

	router.InitRouter(r)

	r.StaticFS("/assets", gin.Dir("./assets", true))

	err := r.Run()
	if err != nil {
		return
	}
}
