package main

import (
	"fmt"

	"github.com/CodeRunRepeat/presidio-go-client/client"
)

func main() {
	//var analyzerBaseUrl string = "https://presidio-analyzer-prod.azurewebsites.net"
	var anonymizerBaseUrl string = "https://presidio-anonymizer-prod.azurewebsites.net"

	presidioAsAnEncryptionService(anonymizerBaseUrl)
	presidioAsAHashingService(anonymizerBaseUrl)
}

func presidioAsSomethingElse(analyzerBaseUrl string, anonymizer client.Anonymizer, description string) {
	var c = client.NewClient(analyzerBaseUrl, nil)

	const SENTENCE string = "This is a test"
	anonymizers := make(client.AnonymizerSet)
	anonymizers.AddDefaultAnonymizer(anonymizer)

	var fakeResult client.AnalyzerResult = *client.NewAnalyzerResult(1)
	fakeResult.Matches[0] = client.AnalyzerMatch{Start: 0, End: int32(len(SENTENCE)), Score: 1.00, EntityType: "ANY"}

	transformed, _, err := c.Anonymize(SENTENCE, &anonymizers, &fakeResult)
	printError(err, "Anonymize")

	if err == nil {
		fmt.Printf("%v value of '%v' is '%v'\n", description, SENTENCE, transformed)
	}
}

func presidioAsAHashingService(analyzerBaseUrl string) {
	presidioAsSomethingElse(analyzerBaseUrl, client.HashAnonymizer{HashType: "sha256"}, "Hashed")
}

func presidioAsAnEncryptionService(analyzerBaseUrl string) {
	presidioAsSomethingElse(analyzerBaseUrl, client.EncryptAnonymizer{Key: "AAECAwQFBgcICQoLDA0ODw=="}, "Encrypted")
}

func printError(err error, funcName string) {
	if err != nil {
		fmt.Printf("Error in calling %v(): %v\n", funcName, err)
	}
}
