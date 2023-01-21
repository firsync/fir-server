package main

import "os"

func folderExists(folder string) (bool, error) {
	fileInfo, err := os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func createFile(file string) error {
	_, err := os.Create(file)
	if err != nil {
		return err
	}
	return nil
}

func createFolder(folder string) error {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
