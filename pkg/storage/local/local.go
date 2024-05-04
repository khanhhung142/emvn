package local

import (
	"log/slog"
	"os"
	"path/filepath"
)

type localStorage struct {
	Directory string
}

var localStorageClient localStorage

func InitLocalStorage() {
	// Store file in the /storage/local/files directory
	// The directory will be created if it does not exist
	_ = os.MkdirAll(filepath.Join("storage"), 0755)
	// Get the current working directory
	pwd, _ := os.Getwd()
	localStorageClient = localStorage{Directory: filepath.Join(pwd, "storage")}
}

func Storage() localStorage {
	return localStorageClient
}

func (l localStorage) SaveFile(file []byte, fileName string) (string, error) {
	filePath := l.Directory + "/" + fileName
	emptyFile, err := os.Create(filePath)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}
	defer emptyFile.Close()
	_, err = emptyFile.Write(file)
	// err = os.WriteFile(dir, file, 0644)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}
	return filePath, nil
}

func (l localStorage) GetFile(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (l localStorage) DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
