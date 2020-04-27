//
// Copyright (c) 2020
// All rights reserved
// filename: handler.go
// description: 工具类函数
// version: 0.1.0
// created by micor_JF(micor5688@163.com) at 2020-4-27
//

package handler

import (
	"FileStoreServer/meta"
	"FileStoreServer/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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
		// 初始化文件元信息
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2016-01-02 15:04:05"),
		}
		//创建文件存储位置
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create new file,err: %s\n", err.Error())
			return
		}
		defer newFile.Close()
		//将内存文件中的内容拷贝到新的文件中
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into new file,err: %s\n", err.Error())
			return
		}
		// 计算上传文件的哈希值
		newFile.Seek(0, 0)
		fileMeta.FileShal = utils.FileSha1(newFile)
		//更新文件元信息
		meta.UploadFileMeta(fileMeta)
		//如果上传成功，将html页面重定向为上传成功页面
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucceceHandler 显示上传文件成功信息
func UploadSucceceHandler(w http.ResponseWriter, r *http.Request) {
	//返回文件上传成功页面
	io.WriteString(w, "Upload Finished!")
}

// GetFileMetaHandler 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                             //解析form
	filehash := r.Form["filehash"][0]         //获取文件对应的哈希值
	fileMetaMes := meta.GetFileMeta(filehash) //将文件哈希值存储到文件元信息当中
	// 将元信息进行序列化操作
	data, err := json.Marshal(fileMetaMes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 写入文件信息
	w.Write(data)
}
