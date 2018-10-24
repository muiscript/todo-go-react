package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/todos", func(ctx *gin.Context) {
		var todos []Todo
		db.Find(&todos)

		ctx.JSON(http.StatusOK, todos)
	})

	router.POST("/todos", func(ctx *gin.Context) {
		var todo Todo

		if err := ctx.ShouldBind(&todo); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&todo)

		ctx.JSON(http.StatusOK, todo)
	})

	router.Run(":4000")
}