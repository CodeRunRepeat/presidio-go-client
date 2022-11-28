package client

import (
	"context"

	"presidio.org/generated"
)

type client struct {
	apiClient *generated.APIClient
}

func NewClient(baseUrl string) *client {
	conf := generated.NewConfiguration()
	conf.BasePath = baseUrl
	conf.AddDefaultHeader("Accept", "application/json")

	c := new(client)
	c.apiClient = generated.NewAPIClient(conf)
	return c
}

func (c client) AnalyzeWithDefaults(text string, language string) (AnalyzerResult, error) {
	request := new(generated.AnalyzeRequest)
	request.Text = text
	request.Language = language

	result, _, err := c.apiClient.AnalyzerApi.AnalyzePost(createContext(), *request)
	return transformResult(result), err
}

func (c client) AnalyzeWithPattern(text string, language string, pattern string, threshold float64, entityName string) (AnalyzerResult, error) {
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

	return analyzerResult
}

func createContext() context.Context {
	return context.TODO()
}
