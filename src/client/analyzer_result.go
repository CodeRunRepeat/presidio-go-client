package client

import (
	"fmt"
	"strings"
)

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
}

// String returns a string representation of an AnalyzerMatch.
func (m *AnalyzerMatch) String() string {
	return fmt.Sprintf("%v-%v (%v, %v)", m.Start, m.End, m.Score, m.EntityType)
}

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
