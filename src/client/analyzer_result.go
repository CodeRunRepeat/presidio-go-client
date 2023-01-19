package client

import (
	"fmt"
	"strings"

	"github.com/CodeRunRepeat/presidio-go-client/generated"
)

/*-------------------------------- AnalyzerMatch --------------------------------*/
// AnalyzerMatch represents a single PII entity identifed in text.
type AnalyzerMatch struct {
	// Start is the index of the match's first character in the text.
	Start int32

	// End is the index of the match's end (first non-character) in the text.
	End int32

	// Score is the confidence score of this match.
	Score float64

	// EntityType is the name of the PII entity identified.
	EntityType string

	// RecognizerName is the name of the recognizer that identified this PII.
	RecognizerName string
}

// String returns a string representation of an AnalyzerMatch.
func (m *AnalyzerMatch) String() string {
	return fmt.Sprintf("%v-%v (%v, %v, %v)", m.Start, m.End, m.Score, m.EntityType, m.RecognizerName)
}

/*-------------------------------- AnalyzerResult --------------------------------*/
// AnalyzerResult represents the overall outcome of PII analysis on a text.
type AnalyzerResult struct {
	// Matches contains all the possible PII matches found in the text,
	// which may overlap with each other.
	Matches []AnalyzerMatch
}

// NewAnalyzerResult stands up an AnalyzerResult that can
// initially accommodate numMatches matches.
func NewAnalyzerResult(numMatches int) *AnalyzerResult {
	var r = new(AnalyzerResult)
	r.Matches = make([]AnalyzerMatch, numMatches)
	return r
}

// String returns a string representation of an AnalyzerResult.
func (r *AnalyzerResult) String() string {
	if r.Matches == nil || len(r.Matches) == 0 {
		return ""
	}

	var matches []string = make([]string, len(r.Matches))
	for index, m := range r.Matches {
		matches[index] = m.String()
	}
	return strings.Join(matches, "\n")
}

func transformToAnalyzerResult(result []generated.RecognizerResultWithAnaysisExplanation) AnalyzerResult {
	var analyzerResult = NewAnalyzerResult(len(result))
	for index, r := range result {
		analyzerResult.Matches[index] = AnalyzerMatch{r.Start, r.End, r.Score, r.EntityType, "N/A"}
		if r.RecognitionMetadata != nil {
			analyzerResult.Matches[index].RecognizerName = r.RecognitionMetadata.RecognizerName
		}
	}

	return *analyzerResult
}

//lint:ignore U1000 Will be fixed when we resolve the anonymizer serialization issue with the Presidio team
func transformFromAnalyzerResult(analyzerResult *AnalyzerResult) []generated.RecognizerResult {
	result := make([]generated.RecognizerResult, len(analyzerResult.Matches))
	for index, m := range analyzerResult.Matches {
		result[index] = generated.RecognizerResult{Start: m.Start, End: m.End, Score: m.Score, EntityType: m.EntityType}
		if m.RecognizerName != "" {
			result[index].RecognitionMetadata = &generated.RecognizedMetadata{RecognizerName: m.RecognizerName}
		}
	}

	return result
}

/*-------------------------------- AnalyzerMatchExplanation --------------------------------*/
// AnalyzerMatchExplanation contains the explanation for a single PII match
type AnalyzerMatchExplanation generated.AnalysisExplanation

func (a *AnalyzerMatchExplanation) String() string {
	return fmt.Sprintf("%+v", *a)
}

/*-------------------------------- AnalyzerResultExplanation --------------------------------*/
// AnalyzerResultExplanation contains the explanation for all PII entities identified in
// the input text
type AnalyzerResultExplanation []AnalyzerMatchExplanation

// NewAnalyzerResultExplanation stands up an AnalyzerResultExplanation that can
// initially accommodate numMatches matches' explanations.
func NewAnalyzerResultExplanation(numMatches int) *AnalyzerResultExplanation {
	newExplanation := make([]AnalyzerMatchExplanation, numMatches)
	cast := AnalyzerResultExplanation(newExplanation)
	return &cast
}

// String returns a string representation of an AnalyzerResultExplanation.
func (r *AnalyzerResultExplanation) String() string {
	if len(*r) == 0 {
		return ""
	}

	var matches []string = make([]string, len(*r))
	for index, m := range *r {
		matches[index] = m.String()
	}
	return strings.Join(matches, "\n")
}

func transformExplanation(result []generated.RecognizerResultWithAnaysisExplanation) AnalyzerResultExplanation {
	var explanation = NewAnalyzerResultExplanation(len(result))

	for index, r := range result {
		m := AnalyzerMatchExplanation(*r.AnalysisExplanation)
		(*explanation)[index] = m
	}

	return *explanation
}
