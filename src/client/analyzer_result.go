package client

import (
	"fmt"
	"strings"
)

type AnalyzerMatch struct {
	Start, End int32
	Score      float64
	EntityType string
}

func (m AnalyzerMatch) String() string {
	return fmt.Sprintf("%v-%v (%v, %v)", m.Start, m.End, m.Score, m.EntityType)
}

type AnalyzerResult struct {
	Matches []AnalyzerMatch
}

func NewAnalyzerResult(numMatches int) AnalyzerResult {
	var r = new(AnalyzerResult)
	r.Matches = make([]AnalyzerMatch, numMatches)
	return *r
}

func (r AnalyzerResult) String() string {
	if r.Matches == nil || len(r.Matches) == 0 {
		return ""
	}

	var matches []string = make([]string, len(r.Matches))
	for index, m := range r.Matches {
		matches[index] = m.String()
	}
	return strings.Join(matches, "\n")
}
