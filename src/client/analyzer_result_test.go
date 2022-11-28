package client

import "testing"

func newAnalyzerMatch() AnalyzerMatch {
	var match = new(AnalyzerMatch)
	return *match
}

func TestNewAnalyzerResult(t *testing.T) {
	numMatches := 3
	r := NewAnalyzerResult(numMatches)
	if len(r.Matches) != numMatches {
		t.Errorf("Requested creation with size %v, result has size %v", numMatches, len(r.Matches))
	}
}

func TestAnalyzerMatchString(t *testing.T) {
	match := newAnalyzerMatch()
	s := match.String()

	if len(s) < 1 {
		t.Errorf("AnalyzerMatch.String() return unexpected value %q", s)
	}
}
