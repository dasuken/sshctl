package sshctl

import (
	"io/ioutil"
	"os"
	"testing"
)

var InitData = `
HOST init
	HostName 54.238.176.202
	User ec2-user
	identityFile /key.pem
`

var InitDataMap = []map[string]string{
	{
		"HOST": "init",
		"HostName": "54.238.176.202",
		"User":"ec2-user",
		"identityFile": "/key.pem",
	},
}

var InitDataConfig = Config{
	Host:                  InitDataMap[0]["HOST"],
	HostName:              InitDataMap[0]["HostName"],
	User:                  InitDataMap[0]["User"],
	IdentityFile:          InitDataMap[0]["identityFile"],
	ProxyCommand:          InitDataMap[0]["ProxyCommand"],
}

func CreateTestConfigFile(t *testing.T) (string, func()) {
	f, err := ioutil.TempFile("", "config")
	if err != nil {
		t.Errorf("failed to crete tmpfile %v", err)
	}

	_, err = f.WriteString(InitData)
	if err != nil {
		t.Errorf("failed to write initdata %v\n", err)
	}

	return f.Name(), func() {
		err := os.Remove(f.Name())
		if err != nil {
			t.Fatal("failed to remove tmpfile", err)
		}
	}
}

