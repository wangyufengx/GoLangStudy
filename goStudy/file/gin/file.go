package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func main() {
	httpRouter := gin.Default()

	httpRouter.POST("/upload", func(ctx *gin.Context) {
		forms, err := ctx.MultipartForm()
		if err != nil {
			fmt.Println("error", err)
		}
		files := forms.File["fileName"]
		for _, v := range files {
			if err := ctx.SaveUploadedFile(v, fmt.Sprintf("%s%s", "./file/", v.Filename)); err != nil {
				fmt.Println("保存文件失败")
			}
		}
	})

	httpRouter.GET("/download", func(ctx *gin.Context) {
		filePath := ctx.Query("url")
		//打开文件
		file, err := os.Open("./" + filePath)
		if err != nil {
			fmt.Println("打开文件错误", err)
		}
		defer file.Close()

		//获取文件的名称
		fileName := path.Base(filePath)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; url="+fileName)
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Cache-Control", "no-cache")

		ctx.File("./" + filePath)
	})

	if err := httpRouter.Run(":9000"); err != nil {
		panic(err)
	}
}
