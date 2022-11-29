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

func TestAnalyzeWithDefaults(t *testing.T) {
	client := setupTest()

	result, err := client.AnalyzeWithDefaults("My name is John Smith", "en")

	if err != nil {
		t.Errorf("Analyze() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("Analyze() returned unexpected response %v", result)
	}
}

func TestAnalyzeWithPattern(t *testing.T) {
	client := setupTest()

	result, err := client.AnalyzeWithPattern("My phone is 123456", "en", "\\d+", 0.80, "SIMPLE_PHONE")

	if err != nil {
		t.Errorf("Analyze() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("Analyze() returned unexpected response %v", result)
	}
}

func TestAnalyzeWithOptions(t *testing.T) {
	client := setupTest()

	options := new(AnalyzerOptions)
	options.
		SetCorrelationId("my_correlation_id").
		AddDenyList("DENY_1", "OTHER_NAME", "en", []string{"Lampros"}, nil)

	result, err := client.AnalyzeWithOptions("My name is Lampros Smith and phone is 123456", "en", options)

	if err != nil {
		t.Errorf("Analyze() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("Analyze() returned unexpected response %v", result)
	}
}
