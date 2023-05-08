package util

import (
	"io"
	"mime/multipart"
	"os"
)

//
// SaveUploadedFile 将前端上传的文件保存到 dst 参数指定目录
//	@parameter	file 前端上传的文件对象
//	@parameter	dst 目标文件, 包含路径和目标文件名
//
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = io.Copy(out, src)
	return err
}
