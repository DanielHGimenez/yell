package config

import "path"

func GetDefaultNotificationSoundFilePath() (string, error) {
	execDir, err := getExecutablePath()
	if err != nil {
		return "", err
	}
	return path.Join(execDir, "notification.mp3"), nil
}
