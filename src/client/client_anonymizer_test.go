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
	t.Skip("Anonymize calls currently fail due to generated client issue.")
	client := setupTest(nil, ANONYMIZER_CLIENT)

	var anonymizers = CreateAnonymizerSet()
	anonymizers.AddDefaultAnonymizer(ReplaceAnonymizer{NewValue: "****"})

	_, err := client.Anonymize("My name is Joe", anonymizers, nil)
	if err != nil {
		t.Error(err.Error())
	}
}
