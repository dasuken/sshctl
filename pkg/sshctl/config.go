package sshctl

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	ErrEmptyConfig       = errors.New("Invalid config")
	ErrEmptyHost         = errors.New("Can't be blank Host")
	ErrEmptyHostName     = errors.New("Can't be blank HostName")
	ErrEmptyUser         = errors.New("Can't be blank User")
	ErrEmptyIdentityFile = errors.New("Can't be blank IdentityFile")
	ErrEmptyPath         = errors.New("Can't be blank config path")
)

type Config struct {
	Command               string
	Host                  string
	HostName              string
	Port                  string
	User                  string
	IdentityFile          string
	ProxyCommand          string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Valid() error {
	if c == nil {
		return ErrEmptyConfig
	}

	if c.Host == "" {
		return ErrEmptyHost
	}

	if c.HostName == "" {
		return ErrEmptyHostName
	}

	if c.User == "" {
		return ErrEmptyUser
	}

	if c.IdentityFile == "" {
		return ErrEmptyIdentityFile
	}

	return nil
}

func (c *Config) DumpFormattedString() string {
	buf := &bytes.Buffer{}

	fmt.Fprintln(buf, "")
	fmt.Fprintf(buf, "HOST %s\n", c.Host)
	fmt.Fprintf(buf, "	HostName %s\n", c.HostName)
	fmt.Fprintf(buf, "	User %s\n", c.User)
	if c.Port != "" {
		fmt.Fprintf(buf, "	Port %s\n", c.Port)
	}
	if c.IdentityFile != "" {
		fmt.Fprintf(buf, "	identityFile %s\n", c.IdentityFile)
	}
	if c.ProxyCommand != "" {
		fmt.Fprintf(buf, "	ProxyCommand %s\n", c.ProxyCommand)
	}

	return buf.String()
}
