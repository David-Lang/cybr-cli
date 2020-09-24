package api_test

import (
	"strings"
	"testing"

	pasapi "github.com/infamousjoeg/pas-api-go/pkg/cybr/api"
)

func TestCyberarkLogonSuccess(t *testing.T) {
	client := pasapi.Client{
		Hostname: hostname,
		AuthType: "cyberark",
	}

	creds := pasapi.LogonRequest{
		Username: username,
		Password: password,
	}

	err := client.Logon(creds)
	if err != nil {
		t.Errorf("Failed to logon. %s", err)
	}
}

func TestCyberarkLogonInvalidCreds(t *testing.T) {
	client := pasapi.Client{
		Hostname: hostname,
		AuthType: "cyberark",
	}

	creds := pasapi.LogonRequest{
		Username: "notReal",
		Password: password,
	}

	err := client.Logon(creds)
	if err == nil {
		t.Errorf("Successfully logged in but shouldn't have. %s", err)
	}
}

func TestCyberarkLogonInvalidHostName(t *testing.T) {
	client := pasapi.Client{
		Hostname: "https://invalidhostname",
		AuthType: "cyberark",
	}

	creds := pasapi.LogonRequest{
		Username: "notReal",
		Password: password,
	}

	err := client.Logon(creds)
	if err == nil {
		t.Errorf("Successfully logged in but shouldn't have. %s", err)
	}
}

func TestLogonInvalidAuthType(t *testing.T) {
	client := pasapi.Client{
		Hostname: hostname,
		AuthType: "notGood",
	}

	creds := pasapi.LogonRequest{
		Username: username,
		Password: password,
	}

	err := client.Logon(creds)
	if err == nil {
		t.Errorf("Successfully logged in but shouldn't have. %s", err)
	}

	if !strings.Contains(err.Error(), "Invalid auth type 'notGood'") {
		t.Errorf("Recieved incorrect error message. %s", err)
	}
}

func TestCyberarkLogoffSuccess(t *testing.T) {
	client := pasapi.Client{
		Hostname: hostname,
		AuthType: "cyberark",
	}

	creds := pasapi.LogonRequest{
		Username: username,
		Password: password,
	}

	err := client.Logon(creds)
	if err != nil {
		t.Errorf("Failed to logon. %s", err)
	}

	err = client.Logoff()
	if err != nil {
		t.Errorf("Failed to logoff. %s", err)
	}
}

func TestCyberarkLogoffFailNotLoggedIn(t *testing.T) {
	client := pasapi.Client{
		Hostname: hostname,
		AuthType: "cyberark",
	}

	err := client.Logoff()
	if !strings.Contains(err.Error(), "401") {
		t.Errorf("Expected to recieve 401 statuc code. %s", err)
	}
}
