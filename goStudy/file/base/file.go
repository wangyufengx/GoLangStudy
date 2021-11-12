package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

func main() {
	server := http.Server{
		Addr: ":9000",
	}

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(20); err != nil {
			fmt.Println("解析url失败", err)
		}
		if r.Method == "GET" {
			files, err := template.ParseFiles("fileName")
			if err != nil {
				fmt.Println("解析文件失败", err)
			}
			files.Execute(w, nil)
		} else {
			files := r.MultipartForm.File["fileName"]
			for _, v := range files {
				f, err := os.Create("./file/" + v.Filename)
				if err != nil {
					fmt.Println("创建文件失败")
				}
				file, err := v.Open()
				if err != nil {
					fmt.Println("打开文件失败")
				}
				_, err = io.Copy(f, file)
				if err != nil {
					fmt.Println("拷贝文件数据失败")
				}
			}
			w.Write([]byte("upload success"))
		}
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println("解析url失败", err)
		}
		fileName := r.Form["fileName"][0]
		filePath := "./" + fileName
		_, err := os.Stat(filePath)
		if err != nil || os.IsNotExist(err) {
			fmt.Println("文件不存在", err)
		}
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("读取文件失败", err)
		}
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Disposition", "attachment; filename= "+fileName)
		w.Write(bytes)
	})
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
