package client

import (
	"golang.org/x/oauth2"
	"presidio.org/generated"
)

type AuthenticationMethod interface {
	getAuthenticationMethod() AuthenticationMethod
}

type BasicAuth generated.BasicAuth

func (b *BasicAuth) getAuthenticationMethod() AuthenticationMethod { return b }

type APIKey generated.APIKey

func (a *APIKey) getAuthenticationMethod() AuthenticationMethod { return a }

type OAuthToken string

func (o *OAuthToken) getAuthenticationMethod() AuthenticationMethod { return o }

type TokenSource struct{ TokenSource oauth2.TokenSource }

func (t *TokenSource) getAuthenticationMethod() AuthenticationMethod { return t }
