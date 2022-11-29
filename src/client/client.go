/*
Client in Golang for the Presidio tool https://github.com/microsoft/presidio.
Presidio is a context aware, pluggable and customizable PII anonymization service for text and images.

This client is not an official part of Presidio.
*/
package client

import (
	"context"

	"presidio.org/generated"
)

// A Client represents a Presidio client, used to call Presidio services that analyzes and anonymizes PII.
type Client struct {
	apiClient *generated.APIClient
}

// NewClient creates a new client to a service located at baseUrl
func NewClient(baseUrl string) *Client {
	conf := generated.NewConfiguration()
	conf.BasePath = baseUrl
	conf.AddDefaultHeader("Accept", "application/json")

	c := new(Client)
	c.apiClient = generated.NewAPIClient(conf)
	return c
}

// AnalyzeWithDefaults analyzes text for PII in a specific language, using the default configuration,
// and returns an AnalyzerResult containing the entities found.
func (c *Client) AnalyzeWithDefaults(text string, language string) (AnalyzerResult, error) {
	request := new(generated.AnalyzeRequest)
	request.Text = text
	request.Language = language

	result, _, err := c.apiClient.AnalyzerApi.AnalyzePost(createContext(), *request)
	return transformResult(result), err
}

// AnalyzeWithPattern analyzes text for PII in a specific language, including a regex based custom entity called entityName,
// with a specified score threshold,
// and returns an AnalyzerResult containing the entities found.
func (c *Client) AnalyzeWithPattern(text string, language string, pattern string, threshold float64, entityName string) (AnalyzerResult, error) {
	request := new(generated.AnalyzeRequest)
	request.Text = text
	request.Language = language

	request.AdHocRecognizers = make([]generated.PatternRecognizer, 1)
	request.AdHocRecognizers[0].Name = "CUSTOM_" + entityName + "_" + language
	request.AdHocRecognizers[0].SupportedEntity = entityName
	request.AdHocRecognizers[0].SupportedLanguage = language
	request.AdHocRecognizers[0].Patterns = make([]generated.Pattern, 1)
	request.AdHocRecognizers[0].Patterns[0].Name = request.AdHocRecognizers[0].Name
	request.AdHocRecognizers[0].Patterns[0].Regex = pattern
	request.AdHocRecognizers[0].Patterns[0].Score = threshold

	result, _, err := c.apiClient.AnalyzerApi.AnalyzePost(createContext(), *request)
	return transformResult(result), err
}

func transformResult(result []generated.RecognizerResultWithAnaysisExplanation) AnalyzerResult {
	var analyzerResult = NewAnalyzerResult(len(result))
	for index, r := range result {
		m := AnalyzerMatch{r.Start, r.End, r.Score, r.EntityType}
		analyzerResult.Matches[index] = m
	}

	return *analyzerResult
}

func createContext() context.Context {
	return context.TODO()
}
