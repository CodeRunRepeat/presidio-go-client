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

	result, _, err := c.apiClient.AnalyzerApi.AnalyzePost(context.TODO(), *request)

	var analyzerResult = NewAnalyzerResult(len(result))
	for index, r := range result {
		m := AnalyzerMatch{r.Start, r.End, r.Score, r.EntityType}
		analyzerResult.Matches[index] = m
	}

	return analyzerResult, err
}

func (c client) AnalyzeWithPattern(text string, language string, pattern string, threshold float64) {

}
