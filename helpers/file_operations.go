package helpers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func FileNameWOExtension(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func ReadFile(filePath string) ([]string, error) {
	_, errExists := os.Stat(filePath)
	if errExists != nil {
		return nil, fmt.Errorf("not a valid file parsed")

	}

	fileHandler, errOp := os.Open(filePath)
	if errOp != nil {
		return nil, errOp
	}

	var errClo error
	defer func() {
		errClo = fileHandler.Close()
	}()

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("empty file passed")
	}

	return res, errClo
}
