package client

import "testing"

func TestAnonymizerHealth(t *testing.T) {
	client := setupTest(nil, ANONYMIZER_CLIENT)

	_, err := client.AnonymizerHealth()
	if err != nil {
		t.Errorf("AnonymizerHealth() failed with error %q", err)
	}
}

func TestAnonymize(t *testing.T) {
	client := setupTest(nil, ANONYMIZER_CLIENT)

	var anonymizers = CreateAnonymizerSet()
	anonymizers.AddDefaultAnonymizer(ReplaceAnonymizer{NewValue: "****"})

	var anResult AnalyzerResult
	anResult.Matches = make([]AnalyzerMatch, 1)
	anResult.Matches[0] = AnalyzerMatch{Start: 11, End: 21, Score: 0.85, EntityType: "PERSON"}

	_, _, err := client.Anonymize("My name is John Smith", anonymizers, &anResult)
	if err != nil {
		t.Errorf("Anonymize() failed with error %q", err)
	}
}

func TestDeanonymize(t *testing.T) {
	client := setupTest(nil, ANONYMIZER_CLIENT)

	var anonymizers = CreateAnonymizerSet()
	anonymizers.AddDefaultAnonymizer(EncryptAnonymizer{Key: "AAECAwQFBgcICQoLDA0ODw=="})

	var anResult AnonymizerResult
	anResult.Matches = make([]AnonymizerMatch, 1)
	anResult.Matches[0] = AnonymizerMatch{Start: 11, End: 55, EntityType: "PERSON", AnonymizerAction: "encrypt", Text: "t99qIN/lfJICokb9voMr4TnxhqLGMR5ubpZrl6oqqJc="}
	newText, _, err := client.Deanonymize("My name is t99qIN/lfJICokb9voMr4TnxhqLGMR5ubpZrl6oqqJc=", anonymizers, &anResult)
	if err != nil {
		t.Errorf("Deanonymize() failed with error %q", err)
	}

	originalText := "My name is John Smith"
	if newText != originalText {
		t.Errorf("Deanonymize() failed: expected '%v', got '%v'", originalText, newText)
	}
}

func TestGetAnonymizers(t *testing.T) {
	client := setupTest(nil, ANONYMIZER_CLIENT)

	anonymizers, err := client.GetAnonymizers()
	if err != nil {
		t.Errorf("GetAnonymizers() failed with error %q", err)
	}
	if len(anonymizers) < 5 {
		t.Errorf("GetAnonymizers() expected to return at least 5 values, returned %v", len(anonymizers))
	}
}

func TestGetDeanonymizers(t *testing.T) {
	client := setupTest(nil, ANONYMIZER_CLIENT)

	deanonymizers, err := client.GetDeanonymizers()
	if err != nil {
		t.Errorf("GetDeanonymizers() failed with error %q", err)
	}
	if len(deanonymizers) < 1 {
		t.Errorf("GetDeanonymizers() expected to return at least 1 value, returned %v", len(deanonymizers))
	}
}
