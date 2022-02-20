package sshctl

import (
	"errors"
	"fmt"
	"os"
)

type Client struct {
	Path          string
	ContentLength string
}

func NewClient(src string) *Client {
	return &Client{
		Path:          src,
	}
}

func (c *Client) ReadAll() ([]Config, error) {
	if c.Path == "" {
		return nil, errors.New("path must exist")
	}

	f, err := os.OpenFile(c.Path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	// ssh configをデコードするか？
	list, err := Decode(f)
	if err != nil {
		return nil, err
	}

	var configList []Config
	for _, v := range list {
		config := Config{}

		config.Host = v["HOST"]
		config.HostName = v["HostName"]
		config.ProxyCommand = v["ProxyCommand"]
		config.User = v["User"]
		config.IdentityFile = v["identityFile"]

		if err := config.Valid(); err != nil {
			return nil, err
		}

		configList = append(configList, config)
	}

	return configList, nil
}

func (c *Client) List() ([]string, error) {
	if c.Path == "" {
		return nil, errors.New("path must exist")
	}

	f, err := os.OpenFile(c.Path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	// ssh configをデコードするか？
	list, err := Decode(f)
	if err != nil {
		return nil, err
	}

	hosts := make([]string, len(list))
	for i, v := range list {
		hosts[i] = v["HOST"]
	}

	return hosts, nil
}


func (c *Client) Put(config *Config) (int, error) {
	err := config.Valid()
	if err != nil {
		return -1, err
	}

	f, err := os.OpenFile(c.Path, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	data := config.DumpFormattedString()

	return f.WriteString(data)
}

/*
一覧撮ってきて_configに書き込むのが良いかな
データ消えるの怖いからbackupとして残すのもいいかもね
*/
func (c *Client) Update(index int, config Config) error {
	f, err := os.OpenFile(c.Path, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
	}
	defer f.Close()

	configList, err := c.ReadAll()
	if err != nil {
		return err
	}

	if len(configList) == 0 {
		return errors.New("can't find any configs")
	}

	if len(configList) < index {
		return fmt.Errorf("%d is out of length", index)
	}

	return nil
}