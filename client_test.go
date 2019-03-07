package xrp

import (
	"testing"
)

func TestDial(t *testing.T) {

	_, err := Dial("s.altnet.rippletest.net:51233")

	if err != nil {
		t.Error("dial err:", err)
	}

	t.Log("dial success")
}
