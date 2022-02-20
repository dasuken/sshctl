package sshctl

import (
	"testing"
)

func TestClient_Write(t *testing.T) {
	path, fn := CreateTestConfigFile(t)
	defer fn()

	client := NewClient(path)
	var config = &Config{
		Host:                  "test",
		HostName:              "54.238.176.202",
		User:                  "ec2-user",
		IdentityFile:          "/identify.pem",
	}

	out, err :=	client.Put(config)
	if err != nil {
		t.Error(err)
	}
	expected := len(config.DumpFormattedString())

	if expected != out {
		t.Error("don't match expected value")
	}
}

func TestClient_ReadAll(t *testing.T) {
	path, fn := CreateTestConfigFile(t)
	defer fn()

	client := NewClient(path)
	configs, err := client.ReadAll()
	if err != nil {
		t.Error(err)
	}

	if len(configs) > 1 {
		t.Error("test file must exists only 1 element")
	}

	// 後から
	if configs[0].Host != InitDataMap[0]["HOST"] {
		t.Errorf("Host must same name, got %v, expected: %v", configs[0].Host, InitDataMap[0]["HOST"])
	}
}

func TestClient_List(t *testing.T) {
	path, fn := CreateTestConfigFile(t)
	defer fn()

	client := NewClient(path)
	hostList, err := client.List()
	if err != nil {
		t.Error(err)
	}

	if len(hostList) > 1 {
		t.Error("hostList must exist only 1")
	}

	if hostList[0] != InitDataMap[0]["HOST"] {
		t.Errorf("got: %v, expected: %v", hostList[0], InitDataMap[0]["HOST"])
	}
}