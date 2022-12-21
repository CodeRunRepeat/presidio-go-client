/*
Client in Golang for the Presidio tool https://github.com/microsoft/presidio.
Presidio is a context aware, pluggable and customizable PII anonymization service for text and images.

This client is not an official part of Presidio.
*/
package client

import (
	"context"

	"github.com/CodeRunRepeat/presidio-go-client/generated"

	"github.com/antihax/optional"
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
	var options AnalyzerOptions

	options.AddPattern(
		"CUSTOM_"+entityName+"_"+language,
		entityName,
		language,
		[]string{pattern},
		[]float64{threshold},
		nil)

	return c.AnalyzeWithOptions(text, language, &options)
}

// AnalyzeWithOptions analyzes text for PII in a specific language and returns an AnalyzerResult containing the entities found.
// Additional analyzer configuration is provided with the options parameter, which is of type AnalyzerOptions.
func (c *Client) AnalyzeWithOptions(text string, language string, options *AnalyzerOptions) (AnalyzerResult, error) {
	if options == nil {
		return c.AnalyzeWithDefaults(text, language)
	}

	request := new(generated.AnalyzeRequest)
	*request = (options.request) // Shallow copy, intentional

	request.Text = text
	request.Language = language

	result, _, err := c.apiClient.AnalyzerApi.AnalyzePost(createContext(), *request)
	return transformResult(result), err
}

func (c *Client) ExplainWithOptions(text string, language string, options *AnalyzerOptions) (AnalyzerResult, AnalyzerResultExplanation, error) {
	if options == nil {
		var emptyOptions AnalyzerOptions
		options = &emptyOptions
	}

	request := new(generated.AnalyzeRequest)
	*request = (options.request) // Shallow copy, intentional

	request.Text = text
	request.Language = language
	request.ReturnDecisionProcess = true

	result, _, err := c.apiClient.AnalyzerApi.AnalyzePost(createContext(), *request)
	return transformResult(result), transformExplanation(result), err
}

// Health checks the health status of the service and returns a value that indicates success.
func (c *Client) Health() (string, error) {
	result, _, err := c.apiClient.AnalyzerApi.HealthGet(createContext())
	return result, err
}

// Recognizers returns the recognizers supported by the analyzer service. If language is an empty string,
// recognizers for all languages are returned, otherwise the function returns recognizers for the
// language specified.
func (c *Client) Recognizers(language string) ([]string, error) {
	var options generated.AnalyzerApiRecognizersGetOpts
	if language == "" {
		options.Language = optional.EmptyString()
	} else {
		options.Language = optional.NewString(language)
	}

	result, _, err := c.apiClient.AnalyzerApi.RecognizersGet(createContext(), &options)
	return result, err
}

// SupportedEntities returns the PII entities supported by the analyzer service. If language is an empty string,
// entities for all languages are returned, otherwise the function returns entities for the
// language specified.
func (c *Client) SupportedEntities(language string) ([]string, error) {
	var options generated.AnalyzerApiSupportedentitiesGetOpts
	if language == "" {
		options.Language = optional.EmptyString()
	} else {
		options.Language = optional.NewString(language)
	}

	result, _, err := c.apiClient.AnalyzerApi.SupportedentitiesGet(createContext(), &options)
	return result, err
}

/* -------------------- Private methods and functions -------------------- */

func transformResult(result []generated.RecognizerResultWithAnaysisExplanation) AnalyzerResult {
	var analyzerResult = NewAnalyzerResult(len(result))
	for index, r := range result {
		m := AnalyzerMatch{r.Start, r.End, r.Score, r.EntityType, "N/A"}
		if r.RecognitionMetadata != nil {
			m.RecognizerName = r.RecognitionMetadata.RecognizerName
		}
		analyzerResult.Matches[index] = m
	}

	return *analyzerResult
}

func transformExplanation(result []generated.RecognizerResultWithAnaysisExplanation) AnalyzerResultExplanation {
	var explanation = NewAnalyzerResultExplanation(len(result))

	for index, r := range result {
		m := AnalyzerMatchExplanation(*r.AnalysisExplanation)
		(*explanation)[index] = m
	}

	return *explanation
}

func createContext() context.Context {
	return context.TODO()
}
