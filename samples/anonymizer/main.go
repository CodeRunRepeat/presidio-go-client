package main

import (
	"fmt"

	"github.com/CodeRunRepeat/presidio-go-client/client"
)

const (
	ENCRYPTION_KEY = "AAECAwQFBgcICQoLDA0ODw=="
	HASH_ALGORITHM = "sha256"
)

func main() {
	var analyzerBaseUrl string = "https://presidio-analyzer-prod.azurewebsites.net"
	var anonymizerBaseUrl string = "https://presidio-anonymizer-prod.azurewebsites.net"

	encryptSample(analyzerBaseUrl, anonymizerBaseUrl)
	hashSample(analyzerBaseUrl, anonymizerBaseUrl)
	maskSample(analyzerBaseUrl, anonymizerBaseUrl)
	replaceSample(analyzerBaseUrl, anonymizerBaseUrl)
	redactSample(analyzerBaseUrl, anonymizerBaseUrl)
	presidioAsAnEncryptionService(anonymizerBaseUrl)
	presidioAsAHashingService(anonymizerBaseUrl)
}

func encryptSample(analyzerBaseUrl string, anonymizerBaseUrl string) {
	anonymizeSample(analyzerBaseUrl, anonymizerBaseUrl, client.EncryptAnonymizer{Key: ENCRYPTION_KEY})
}

func hashSample(analyzerBaseUrl string, anonymizerBaseUrl string) {
	anonymizeSample(analyzerBaseUrl, anonymizerBaseUrl, client.HashAnonymizer{HashType: HASH_ALGORITHM})
}

func maskSample(analyzerBaseUrl string, anonymizerBaseUrl string) {
	anonymizeSample(analyzerBaseUrl, anonymizerBaseUrl, client.MaskAnonymizer{MaskingChar: "*", CharsToMask: 8})
}

func redactSample(analyzerBaseUrl string, anonymizerBaseUrl string) {
	anonymizeSample(analyzerBaseUrl, anonymizerBaseUrl, client.RedactAnonymizer{})
}

func replaceSample(analyzerBaseUrl string, anonymizerBaseUrl string) {
	anonymizeSample(analyzerBaseUrl, anonymizerBaseUrl, client.ReplaceAnonymizer{NewValue: "<REMOVED>"})
}

func anonymizeSample(analyzerBaseUrl string, anonymizerBaseUrl string, anonymizer client.Anonymizer) {
	var analyzerClient = client.NewClient(analyzerBaseUrl, nil)
	const SENTENCE string = "On September 18 I visited microsoft.com and sent an email to test@presidio.site, from the IP 192.168.0.1"
	analyzerResult, err := analyzerClient.AnalyzeWithDefaults(SENTENCE, "en")

	if printError(err, "AnalyzeWithDefaults") {
		return
	}

	var anonymizerClient = client.NewClient(anonymizerBaseUrl, nil)
	anonymizers := make(client.AnonymizerSet)
	anonymizers.AddDefaultAnonymizer(anonymizer)
	anonymizedText, _, err := anonymizerClient.Anonymize(SENTENCE, &anonymizers, &analyzerResult)

	if !printError(err, "Anonymize") {
		fmt.Printf("Anonymized text: '%v'\n", anonymizedText)
	}
}

func presidioAsSomethingElse(anonymizerBaseUrl string, anonymizer client.Anonymizer, description string) {
	var c = client.NewClient(anonymizerBaseUrl, nil)

	const SENTENCE string = "This is a test"
	anonymizers := make(client.AnonymizerSet)
	anonymizers.AddDefaultAnonymizer(anonymizer)

	var fakeResult client.AnalyzerResult = *client.NewAnalyzerResult(1)
	fakeResult.Matches[0] = client.AnalyzerMatch{Start: 0, End: int32(len(SENTENCE)), Score: 1.00, EntityType: "ANY"}

	transformed, _, err := c.Anonymize(SENTENCE, &anonymizers, &fakeResult)

	if !printError(err, "Anonymize") {
		fmt.Printf("%v value of '%v' is '%v'\n", description, SENTENCE, transformed)
	}
}

func presidioAsAHashingService(anonymizerBaseUrl string) {
	presidioAsSomethingElse(anonymizerBaseUrl, client.HashAnonymizer{HashType: HASH_ALGORITHM}, "Hashed")
}

func presidioAsAnEncryptionService(anonymizerBaseUrl string) {
	presidioAsSomethingElse(anonymizerBaseUrl, client.EncryptAnonymizer{Key: ENCRYPTION_KEY}, "Encrypted")
}

func printError(err error, funcName string) bool {
	if err != nil {
		fmt.Printf("Error in calling %v(): %v\n", funcName, err)
		return true
	}
	return false
}
