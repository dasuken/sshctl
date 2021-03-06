package cmd

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/dasuken/sshctl/pkg/flagstr"
	"github.com/dasuken/sshctl/pkg/sshctl"
	"github.com/urfave/cli/v2"
	"strings"
)

func add(ctx *cli.Context) error {
	configPath, err := getDefaultConfigPath()
	if err != nil {
		return err
	}

	// select answer format type
	var ans string
	if err := survey.Ask(sshctl.MakeChoiceAddQuestion(), &ans); err != nil {
		return err
	}

	prefix := ans[:1]
	config, err := getConfig(prefix)

	client := sshctl.NewClient(configPath)
	_, err = client.Put(config)
	if err != nil {
		return err
	}

	sshctl.ShowMessage(
		"success",
		fmt.Sprintf("config was created!! if you use that config setting, $ ssh %s", config.Host),
		true, false,
	)

	sshctl.ShowMessage(
		"",
		fmt.Sprint("*if you want to use more options, please write directory."),
		true, false,
	)

	return nil
}

func getConfig(prefix string) (*sshctl.Config, error) {
	config := &sshctl.Config{}

	// case(prefix == 0) make config from string like 'ssh -i xxx.pem user@hostname'
	// case(prefix == 1) make config f rom interactive
	switch prefix {
	case "0":
		type addAnswer struct {
			Command, Host string
		}
		a := addAnswer{}

		err := survey.Ask(sshctl.MakeAddQuestionByCommandLine(), &a)
		if err != nil {
			return nil, err
		}

		config, err = mapping(a.Command)
		if err != nil {
			return nil, err
		}

		config.Host = a.Host
	case "1":
		err := survey.Ask(sshctl.MakeAddQuestionByInteractive(), config)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid question")
	}

	return config, nil
}

func mapping(str string) (*sshctl.Config, error) {
	splitted := strings.Split(str, " ")
	if splitted[0] != "ssh" {
		return nil, errors.New("please input ssh command")
	}

	userAndHostname := strings.Split(splitted[len(splitted)-1], "@")
	if len(userAndHostname) != 2 {
		return nil, errors.New("Please input 'username@hostname' after any options")
	}

	config := &sshctl.Config{}
	config.User 	= userAndHostname[0]
	config.HostName = userAndHostname[1]

	f := flagstr.New(str)
	config.IdentityFile = f.Get("i")
	config.Port		    = f.Get("p")

	return config, nil
}



