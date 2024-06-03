package main

import (
	"os"
)

func readCSVFromFolder(folderPath, fileName string) (*os.File, error) {
	filePath := folderPath + "/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
