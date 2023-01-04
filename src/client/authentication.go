package client

import (
	"github.com/CodeRunRepeat/presidio-go-client/generated"
	"golang.org/x/oauth2"
)

type AuthenticationMethod interface {
	getAuthenticationMethod() AuthenticationMethod
}

type BasicAuth generated.BasicAuth

func (b BasicAuth) getAuthenticationMethod() AuthenticationMethod { return b }

type APIKey generated.APIKey

func (a APIKey) getAuthenticationMethod() AuthenticationMethod { return a }

type AccessToken string

func (o AccessToken) getAuthenticationMethod() AuthenticationMethod { return o }

type TokenSource struct{ TokenSource oauth2.TokenSource }

func (t TokenSource) getAuthenticationMethod() AuthenticationMethod { return t }
