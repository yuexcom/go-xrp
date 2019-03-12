package xrp

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestDial(t *testing.T) {

	_, err := Dial("s.altnet.rippletest.net:51233", true)

	if err != nil {
		t.Error("dial err:", err)
	}

	t.Log("dial success")
}
