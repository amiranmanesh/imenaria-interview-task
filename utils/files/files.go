package files

import (
	"fmt"
	"github.com/amiranmanesh/imenaria-interview-task/utils/encrypting"
	"mime/multipart"
	"os"
)

type iFiles interface {
	Save(multipart.File, *multipart.FileHeader) (string, error)
}

type filesHandler struct{}

var FilesHandler iFiles = &filesHandler{}

func (filesHandler) Save(file multipart.File, multipartFileHeader *multipart.FileHeader) (string, error) {

	mime, err := GetFileMIME(file)
	if err != nil {
		return "", err
	}
	if mime != "image/png" {
		return "", fmt.Errorf("only image/png files supported")
	}

	src, err := multipartFileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	mBFileSize := multipartFileHeader.Size / 1024 / 1024
	if mBFileSize > 2 {
		return "", fmt.Errorf("file size is too large")
	}

	sha512Hash, err := encrypting.FileSha512(src)
	if err != nil {
		return "", err
	}

	_, err = os.Stat("avatars")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("avatars", 0755)
		if errDir != nil {
			return "", err
		}
	}

	filePath := fmt.Sprintf("avatars/%s.png", sha512Hash)
	if err := SaveUploadedFile(multipartFileHeader, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}
