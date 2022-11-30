package client

import (
	"os"
	"testing"
)

func setupTest() *Client {
	if baseUrl == "" {
		url, found := os.LookupEnv(ENVIROMENT_BASE_URL)
		if found {
			baseUrl = url
		}
	}

	if baseUrl == "" {
		baseUrl = DEFAULT_BASE_URL
	}

	return NewClient(baseUrl)
}

var baseUrl string = ""

const DEFAULT_BASE_URL = "https://presidio-analyzer-prod.azurewebsites.net"
const ENVIROMENT_BASE_URL = "PRESIDIO_CLIENT_BASE_URL"

func TestNewClient(t *testing.T) {
	client := setupTest()

	if client == nil {
		t.Error("NewClient() must return non-nil")
	}
}
