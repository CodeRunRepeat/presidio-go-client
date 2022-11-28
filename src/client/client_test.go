package client

import (
	"os"
	"testing"
)

func setupTest() {
	if baseUrl == "" {
		url, found := os.LookupEnv(ENVIROMENT_BASE_URL)
		if found {
			baseUrl = url
		}
	}

	if baseUrl == "" {
		baseUrl = DEFAULT_BASE_URL
	}
}

var baseUrl string = ""

const DEFAULT_BASE_URL = "https://presidio-analyzer-prod.azurewebsites.net"
const ENVIROMENT_BASE_URL = "PRESIDIO_CLIENT_BASE_URL"

func TestNewClient(t *testing.T) {
	setupTest()

	client := NewClient(baseUrl)

	if client == nil {
		t.Error("NewClient() must return non-nil")
	}
}

func TestAnalyzeWithDefaults(t *testing.T) {
	setupTest()

	client := NewClient(baseUrl)

	result, err := client.AnalyzeWithDefaults("My name is John Smith", "en")

	if err != nil {
		t.Errorf("Analyze() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("Analyze() returned unexpected response %q", result)
	}
}
