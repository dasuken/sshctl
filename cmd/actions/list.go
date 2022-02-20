package actions

import (
	"errors"
	"fmt"
	sshctl2 "github.com/dasuken/sshctl/pkg/sshctl"
	"github.com/urfave/cli/v2"
	"os"
)

func list(ctx *cli.Context) error {
	// create config file
	homePath, ok := os.LookupEnv("HOME")
	if !ok {
		return errors.New("$HOME is not set")
	}

	configPath := fmt.Sprintf("%s/.ssh/config", homePath)
	if err := createIfNotExists(configPath); err != nil {
		return err
	}

	client := sshctl2.NewClient(configPath)

	hosts, err := client.List()
	if err != nil {
		return err
	}

	for i, host := range hosts {
		fmt.Printf("%d: %s \n", i, host)
	}

	return nil
}

