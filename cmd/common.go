package cmd

import (
	"errors"
	"fmt"
	"os"
)

func getDefaultConfigPath() (string, error) {
	homePath, ok := os.LookupEnv("HOME")
	if !ok {
		return "", errors.New("Can't find $HOME environment")
	}

	configPath := fmt.Sprintf("%s/.ssh/config", homePath)
	if _, err := os.Stat(configPath); err != nil {
		return "", fmt.Errorf("config path is not created. Please touch ~/.ssh/config")
		//  or /etc/ssh/config
	}

	return configPath, nil
}

func createIfNotExists(configPath string) error {
	_, err := os.Stat(configPath)
	if err == os.ErrNotExist {
		os.Create(configPath)
	} else if err != nil {
		return err
	}

	return nil
}
