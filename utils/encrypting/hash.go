package encrypting

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"mime/multipart"
)

func FileSha512(src multipart.File) (string, error) {
	hasher := sha512.New()
	if _, err := io.Copy(hasher, src); err != nil {
		return "", err
	}
	sha256Hash := hex.EncodeToString(hasher.Sum(nil))
	return sha256Hash, nil
}
