package client

import (
	"testing"
)

func TestAnalyzeWithDefaults(t *testing.T) {
	client := setupTest()

	result, err := client.AnalyzeWithDefaults("My name is John Smith", "en")

	if err != nil {
		t.Errorf("AnalyzeWithDefaults() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("AnalyzeWithDefaults() returned unexpected response %v", result)
	}
}

func TestAnalyzeWithPattern(t *testing.T) {
	client := setupTest()

	result, err := client.AnalyzeWithPattern("My phone is 123456", "en", "\\d+", 0.80, "SIMPLE_PHONE")

	if err != nil {
		t.Errorf("AnalyzeWithPattern() failed with error %q", err)
	} else if len(result.Matches) != 2 {
		t.Errorf("AnalyzeWithPattern() returned unexpected response %v", result)
	}
}

func TestAnalyzeWithOptions(t *testing.T) {
	client := setupTest()

	options := new(AnalyzerOptions)
	options.
		SetCorrelationId("my_correlation_id").
		AddDenyList("DENY_1", "OTHER_NAME", "en", []string{"Lampros"}, nil)

	result, err := client.AnalyzeWithOptions("My name is Lampros Smith and phone is 123456", "en", options)

	if err != nil {
		t.Errorf("AnalyzeWithOptions() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("AnalyzeWithOptions() returned unexpected response %v", result)
	}
}

func TestExplainWithOptions(t *testing.T) {
	client := setupTest()

	options := new(AnalyzerOptions)
	result, explanations, err := client.ExplainWithOptions("My phone is 123456", "en", options)

	if err != nil {
		t.Errorf("ExplainWithOptions() failed with error %q", err)
	} else if len(result.Matches) < 1 {
		t.Errorf("ExplainWithOptions() returned unexpected response %v", result)
	} else if len(result.Matches) != len(explanations) {
		t.Errorf("ExplainWithOptions() returned %v results but %v matches", result.Matches, explanations)
	}
}

func TestHealth(t *testing.T) {
	client := setupTest()

	_, err := client.Health()
	if err != nil {
		t.Errorf("Health() failed with error %q", err)
	}
}

func TestRecognizers(t *testing.T) {
	client := setupTest()

	_, err := client.Recognizers("")
	if err != nil {
		t.Errorf("Recognizers() for all languages failed with error %q", err)
	}

	_, err = client.Recognizers("en")
	if err != nil {
		t.Errorf("Recognizers() for language en failed with error %q", err)
	}
}

func TestSupportedEntities(t *testing.T) {
	client := setupTest()

	_, err := client.SupportedEntities("")
	if err != nil {
		t.Errorf("SupportedEntities() for all languages failed with error %q", err)
	}

	_, err = client.SupportedEntities("en")
	if err != nil {
		t.Errorf("SupportedEntities() for language en failed with error %q", err)
	}
}
