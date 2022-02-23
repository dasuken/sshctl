package sshctl

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"strings"
)

var (
	sshValidator = func(ans interface{}) error {
		s := strings.Split(ans.(string), " ")
		lastChar := s[len(s) - 1]

		// user@hostname should be last element
		userAndHost := strings.Split(lastChar, "@")
		if len(userAndHost) != 2 {
			return errors.New("Options should be before user@hostname")
		}

		return nil
	}
)

var(
	MakeChoiceAddQuestion func() []*survey.Question = func() []*survey.Question {
		return []*survey.Question{
			{
				Name: "",
				Prompt: &survey.Select{
					Message: "select input format: ",
					Options: []string{
						"0: input from ssh command (ex: ssh -i xxx.pem -p 22 ec2-user@xxx)",
						"1: input from interactive",
					},
				},
				Validate: survey.Required,
			},
		}
	}

	MakeAddQuestionByCommandLine func() []*survey.Question = func() []*survey.Question {
		return []*survey.Question{
			{
				Name: "Command",
				Prompt: &survey.Input{
					Message: "Input ssh command: ",
					Help: "Available options are -p -i",
				},
				Validate: sshValidator,
			},
			{
				Name: "Host",
				Prompt: &survey.Input{
					Message: "unique label as host: ",
				},
				Validate: survey.Required,
			},
		}
	}

	MakeAddQuestionByInteractive func() []*survey.Question = func() []*survey.Question {
		return []*survey.Question{
			{
				Name: "HOST",
				Prompt: &survey.Input{
					Message: "Enter Unique Name to represent specific endpoint:",
				},
				Validate: survey.Required,
			},
			{
				Name: "HostName",
				Prompt: &survey.Input{
					Message: "Enter HostName that represents the connection destination:",
					Help: "Please type ip address for example 35.190.247.0",
				},
				Validate: survey.Required,
			},
			{
				Name: "User",
				Prompt: &survey.Input{
					Message: "Enter user name: ",
					Default:  "ec2-user",
				},
			},
			{
				Name: "IdentityFile",
				Prompt: &survey.Input{
					Message: "Enter IdentityFile path: ",
					Help: "That is key file. ssh -i xxx.　Commonly used extensions are pem or rsa. If you define $IDENTITY_KEY, use it as default value",
					Default: os.Getenv("IDENTITY_KEY"),
				},
				Validate: survey.Required,
			},
			{
				Name: "ProxyCommand",
				Prompt: &survey.Input{
					Message: "Enter ProxyCommand: ",
					Help: "ex) ssh -W %h:%p test",
				},
			},
		}
	}

	MakeUpdateQuestion func(c *Config) []*survey.Question = func(c *Config) []*survey.Question {
		return []*survey.Question{
			{
				Name: "HOST",
				Prompt: &survey.Input{
					Message: "Enter Unique Name to represent specific endpoint:",
					Default: c.Host,
				},
				Validate: survey.Required,
			},
			{
				Name: "HostName",
				Prompt: &survey.Input{
					Message: "Enter HostName that represents the connection destination:",
					Help: "Please type ip address for example 35.190.247.0",
					Default: c.HostName,
				},
				Validate: survey.Required,
			},
			{
				Name: "User",
				Prompt: &survey.Input{
					Message: "Enter user name: ",
					Default: c.User,
				},
			},
			{
				Name: "IdentityFile",
				Prompt: &survey.Input{
					Message: "Enter IdentityFile path: ",
					Help: "That is key file. ssh -i xxx.　Commonly used extensions are pem or rsa. If you define $IDENTITY_KEY, use it as default value",
					Default: c.IdentityFile,
				},
				Validate: survey.Required,
			},
		}
	}
)
