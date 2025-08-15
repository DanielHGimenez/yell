package config

import (
	"os"
	"path"
)

func getExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return path.Dir(ex), nil
}

func GetFilePathInExecutableFolder(fileName string) (string, error) {
	execDir, err := getExecutablePath()
	if err != nil {
		return "", err
	}
	return path.Join(execDir, fileName), nil
}
