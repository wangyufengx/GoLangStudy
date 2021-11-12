package utils

import (
	"bufio"
	"io"
	"os"
)

//GenerateFile 生成镜像包
func GenerateFile(fileName string, data io.ReadCloser) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	if _, err := io.Copy(w, data); err !=nil {
		return err
	}
	w.Flush()
	return nil
}
