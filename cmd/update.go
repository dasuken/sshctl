package cmd

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	sshctl2 "github.com/dasuken/sshctl/pkg/sshctl"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

/*

*/
func update(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return errors.New("please target index number. sshgen update 1")
	}

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

	configList, err := client.ReadAll()
	if err != nil {
		return err
	}

	arg := ctx.Args().First()
	index, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}

	if index > len(configList) {
		return fmt.Errorf("ca't access %d. Please less than %d", index, len(configList))
	}

	if err := survey.Ask(sshctl2.MakeUpdateQuestion(&configList[index]), &configList[index]); err != nil {
		return err
	}

	f, err := os.OpenFile(configPath+"_tmp", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	for _, config := range configList {
		_, err = f.WriteString(fmt.Sprintf("%s", config.DumpFormattedString()))
		if err != nil {
			return err
		}
	}

	if err := os.Remove(configPath); err != nil {
		return err
	}

	os.Rename(configPath+"_tmp", configPath)

	return nil
}
