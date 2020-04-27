//
// Copyright (c) 2020
// All rights reserved
// filename: filemeta.go
// description: 存储元文件信息
// version: 0.1.0
// created by micor_JF(micor5688@163.com) at 2020-4-27
//

package meta

// FileMeta 文件元信息结构体
type FileMeta struct {
	FileShal string //文件的唯一标志
	FileName string //文件名
	FileSize int64  //文件大小
	Location string //文件存储位置
	UploadAt string //时间戳，由时间格式化后的字符串
}

var fileMetas map[string]FileMeta // 存储文件元信息

// 文件元信息的初始化
func init() {
	fileMetas = make(map[string]FileMeta)
}

// UploadFileMeta 更新文件元信息
func UploadFileMeta(fileMark FileMeta) {
	fileMetas[fileMark.FileShal] = fileMark
}

// GetFileMeta  通过sha1获取文件的元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
