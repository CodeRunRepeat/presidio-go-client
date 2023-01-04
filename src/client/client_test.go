package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/CodeRunRepeat/presidio-go-client/generated"
	"golang.org/x/oauth2"
)

func setupTest(auth AuthenticationMethod) *Client {
	if baseUrl == "" {
		url, found := os.LookupEnv(ENVIROMENT_BASE_URL)
		if found {
			baseUrl = url
		}
	}

	if baseUrl == "" {
		baseUrl = DEFAULT_BASE_URL
	}

	return NewClient(baseUrl, auth)
}

var baseUrl string = ""

const DEFAULT_BASE_URL = "https://presidio-analyzer-prod.azurewebsites.net"
const ENVIROMENT_BASE_URL = "PRESIDIO_CLIENT_BASE_URL"

func TestNewClient(t *testing.T) {
	client := setupTest(nil)

	if client == nil {
		t.Error("NewClient() must return non-nil")
	}
}

func TestNewClientBasicAuth(t *testing.T) {
	client := setupTest(createBasicAuth())
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
	client := setupTest(createAccessToken())
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
	client := setupTest(createAPIKey())
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
	client := setupTest(createTokenSource())
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
