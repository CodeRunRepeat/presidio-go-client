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
