package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/CodeRunRepeat/presidio-go-client/generated"
	"golang.org/x/oauth2"
)

func setupTest(auth AuthenticationMethod, clientType int) *Client {
	analyzerBaseUrl = getBaseUrl(analyzerBaseUrl, ENVIROMENT_ANALYZER_URL, DEFAULT_ANALYZER_URL)
	anonymizerBaseUrl = getBaseUrl(anonymizerBaseUrl, ENVIROMENT_ANONYMIZER_URL, DEFAULT_ANONYMIZER_URL)

	if clientType == ANONYMIZER_CLIENT {
		return NewClient(ClientConfig{BaseUrl: anonymizerBaseUrl, AuthenticationMethod: auth})
	}
	return NewClient(ClientConfig{BaseUrl: analyzerBaseUrl, AuthenticationMethod: auth})
}

func getBaseUrl(currentValue string, envName string, defaultValue string) string {
	if currentValue != "" {
		return currentValue
	}

	url, found := os.LookupEnv(envName)
	if found {
		return url
	}

	return defaultValue
}

var analyzerBaseUrl string = ""
var anonymizerBaseUrl string = ""

const (
	ANALYZER_CLIENT = iota + 1
	ANONYMIZER_CLIENT
)

const DEFAULT_ANALYZER_URL = "https://presidio-analyzer-prod.azurewebsites.net"
const ENVIROMENT_ANALYZER_URL = "PRESIDIO_CLIENT_ANALYZER_URL"
const DEFAULT_ANONYMIZER_URL = "https://presidio-anonymizer-prod.azurewebsites.net"
const ENVIROMENT_ANONYMIZER_URL = "PRESIDIO_CLIENT_ANONYMIZER_URL"

func TestNewAnalyzerClient(t *testing.T) {
	client := setupTest(nil, ANALYZER_CLIENT)

	if client == nil {
		t.Error("TestNewAnalyzerClient() must return non-nil")
	}
}

func TestNewAnonymizerClient(t *testing.T) {
	client := setupTest(nil, ANALYZER_CLIENT)

	if client == nil {
		t.Error("TestNewAnonymizerClient() must return non-nil")
	}
}

func TestNewClientBasicAuth(t *testing.T) {
	client := setupTest(createBasicAuth(), ANALYZER_CLIENT)
	if client == nil {
		t.Error("NewClient() must return non-nil")
	}

	context := client.createContext()
	auth := context.Value(generated.ContextBasicAuth)

	if auth == nil {
		t.Error("BasicAuth missing")
	}

	_, ok := auth.(generated.BasicAuth)
	if !ok {
		t.Error("BasicAuth has wrong type")
	}
}

func TestNewClientAccessToken(t *testing.T) {
	client := setupTest(createAccessToken(), ANALYZER_CLIENT)
	if client == nil {
		t.Error("NewClient() must return non-nil")
	}

	context := client.createContext()
	auth := context.Value(generated.ContextAccessToken)

	if auth == nil {
		t.Error("AccessToken missing")
	}

	_, ok := auth.(string)
	if !ok {
		t.Error("AccessToken has wrong type")
	}
}

func TestNewClientAPIKey(t *testing.T) {
	client := setupTest(createAPIKey(), ANALYZER_CLIENT)
	if client == nil {
		t.Error("NewClient() must return non-nil")
	}

	context := client.createContext()
	auth := context.Value(generated.ContextAPIKey)

	if auth == nil {
		t.Error("APIKey missing")
	}

	_, ok := auth.(generated.APIKey)
	if !ok {
		t.Error("APIKey has wrong type")
	}
}

func TestNewClientTokenSource(t *testing.T) {
	client := setupTest(createTokenSource(), ANALYZER_CLIENT)
	if client == nil {
		t.Error("NewClient() must return non-nil")
	}

	context := client.createContext()
	auth := context.Value(generated.ContextOAuth2)

	if auth == nil {
		t.Error("TokenSource missing")
	}

	_, ok := auth.(oauth2.TokenSource)
	if !ok {
		t.Error("TokenSource has wrong type")
	}
}

func createBasicAuth() BasicAuth {
	return BasicAuth{UserName: "test", Password: "test"}
}

func createAccessToken() AccessToken {
	return "dummy token"
}

func createAPIKey() APIKey {
	return APIKey{Key: "key", Prefix: "prefix"}
}

func createTokenSource() TokenSource {
	// Setup some fake oauth2 configuration
	cfg := &oauth2.Config{
		ClientID:     "1234567",
		ClientSecret: "SuperSecret",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://devnull",
			TokenURL: "https://devnull",
		},
		RedirectURL: "https://devnull",
	}

	// and a fake token
	tok := oauth2.Token{
		AccessToken:  "FAKE",
		RefreshToken: "So Fake",
		Expiry:       time.Now().Add(time.Hour * 100000),
		TokenType:    "Bearer",
	}

	// then a fake tokenSource
	context := context.WithValue(context.TODO(), oauth2.HTTPClient, nil)
	return TokenSource{TokenSource: cfg.TokenSource(context, &tok)}
}
