package utils

import (
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// CheckImageFileType 校验图片扩展名
func CheckImageFileType(file *multipart.FileHeader) bool {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	return allowExts[ext]
}

// FileToBytes 将文件流转换为 byte 切片
func FileToBytes(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	return io.ReadAll(src)
}
