package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func GetFileMIME(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		return "", err
	}

	return http.DetectContentType(fileHeader), nil
}

// "/some/path/on/server/" + newFileName as dst
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		fmt.Println(1, err)
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(3, err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
