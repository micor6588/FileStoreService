package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// UploadHandler 文件加载句柄
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}

		io.WriteString(w, string(data))

	} else if r.Method == "POST" {
		//接收文件流,存储到本地目录
		file, head, err := r.FormFile("file") //默认文件名是file
		if err != nil {
			fmt.Printf("Failed to get data,err: %s\n", err.Error())
			return
		}
		defer file.Close()
		//创建文件存储位置
		newFile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create new file,err: %s\n", err.Error())
			return
		}
		defer newFile.Close()
		//将内存文件中的内容拷贝到新的文件中
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into new file,err: %s\n", err.Error())
			return
		}
		//如果上传成功，将html页面重定向为上传成功页面
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucceceHandler 显示上传文件成功信息
func UploadSucceceHandler(w http.ResponseWriter, r *http.Request) {
	//返回文件上传成功页面
	io.WriteString(w, "Upload Finished!")
}
