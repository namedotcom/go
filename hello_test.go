package namecom

import (
	"os"
	"testing"
)

func TestNameCom_HelloFunc(t *testing.T) {
	resp, err := namecom.HelloFunc()
	if err != nil {
		t.Fatal(err)
	}

	if resp.Username != os.Getenv("NAME_COM_USER") {
		t.Fatal("request not equal")
	}

	t.Logf("Hi, greeting form %v, server time is %v", resp.ServerName, resp.ServerTime)
}
