package sshctl

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	path, fn := CreateTestConfigFile(t)
	defer fn()

	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		t.Error(err)
	}

	m, err := Decode(f)
	if err != nil {
		t.Error(err)
	}

	if !DeepEqualMapString(m[0], InitDataMap[0]) {
		t.Errorf("got: %v, expected: %v", m, InitDataMap)
	}
}

func DeepEqualMapString(m1 map[string]string, m2 map[string]string)  bool {
	if m1["HOST"] == m2["HOST"] &&
		m1["HostName"] == m2["HostName"] &&
		m1["User"] == m2["User"]&&
		m1["identityFile"] == m2["identityFile"] { return true }

	return false
}