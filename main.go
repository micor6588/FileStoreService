package main

import (
	"FileStoreServer/handler"
	"fmt"
	"net/http"
)

func main() {
	//文件访问路由
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucceceHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadFileHandler)
	//监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server,err:%s", err.Error())
	}
}
