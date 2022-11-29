package client

import "testing"

func TestNewAnalyzerOptions(t *testing.T) {
	var options AnalyzerOptions

	if options.request.Text != "" {
		t.Error("New AnalyzerOptions request should be zeroed")
	}
}

func TestAnalyzerOptionsSetCorrelationId(t *testing.T) {
	var options AnalyzerOptions

	correlationId := "1"
	options.SetCorrelationId(correlationId)

	if options.request.CorrelationId != correlationId {
		t.Errorf("AnalyzerOptions.SetCorrelationId should set value to %v, is %v", correlationId, options.request.CorrelationId)
	}
}

func TestAnalyzerOptionsAddDenyList(t *testing.T) {
	var options AnalyzerOptions

	options.AddDenyList("DENY", "PII", "en", []string{"one", "two"}, nil)

	if len(options.request.AdHocRecognizers) != 1 {
		t.Errorf("AddDenyList failed, should find an array of length 1, found length %v", len(options.request.AdHocRecognizers))
	}

	options.AddDenyList("DENY2", "PII", "en", []string{"three", "four"}, nil)

	if len(options.request.AdHocRecognizers) != 2 {
		t.Errorf("AddDenyList follow up failed, should find an array of length 2, found length %v", len(options.request.AdHocRecognizers))
	}
}
