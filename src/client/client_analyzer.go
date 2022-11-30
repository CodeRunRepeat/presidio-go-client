package client

import (
	"presidio.org/generated"

	"github.com/antihax/optional"
)

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
