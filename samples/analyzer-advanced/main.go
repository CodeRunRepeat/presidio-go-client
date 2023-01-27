package main

import (
	"fmt"

	"github.com/CodeRunRepeat/presidio-go-client/client"
)

func main() {
	var baseUrl string = "https://presidio-analyzer-prod.azurewebsites.net"

	changeScoreThreshold(baseUrl)

	detectCustomPattern(baseUrl)

	detectSpecificEntities(baseUrl)
}

func printResults(result client.AnalyzerResult, err error, funcName string) {
	if err != nil {
		fmt.Printf("Error in calling %v(): %v\n", funcName, err)
	} else {
		for _, match := range result.Matches {
			fmt.Printf("Match: %v\n", match)
		}
	}
}

func detectCustomPattern(baseUrl string) {
	var c = client.NewClient(baseUrl, nil)
	const PATTERN string = "[A-Z]{2}\\d{12,16}"
	const SENTENCE string = "My bank account is HG1234123412341235123451234"
	result, err := c.AnalyzeWithPattern(SENTENCE, "en", PATTERN, 0.8, "TEST_IBAN")

	fmt.Printf("Analyzing '%v' with pattern '%v'\n", SENTENCE, PATTERN)
	printResults(result, err, "AnalyzeWithPattern")
}

func detectSpecificEntities(baseUrl string) {
	var c = client.NewClient(baseUrl, nil)

	const SENTENCE string = "My name is Joe Smith and my phone is (555)4168123"
	var options client.AnalyzerOptions

	fmt.Printf("Analyzing '%v' for PERSON\n", SENTENCE)
	options.SetEntities([]string{"PERSON"})
	result, err := c.AnalyzeWithOptions(SENTENCE, "en", &options)
	printResults(result, err, "AnalyzeWithOptions")

	fmt.Printf("Analyzing '%v' for PHONE_NUMBER\n", SENTENCE)
	options.SetEntities([]string{"PHONE_NUMBER"})
	result, err = c.AnalyzeWithOptions(SENTENCE, "en", &options)
	printResults(result, err, "AnalyzeWithOptions")
}

func changeScoreThreshold(baseUrl string) {
	var c = client.NewClient(baseUrl, nil)
	const SENTENCE string = "My name is Joe Smith and my phone is (555)4168123"

	fmt.Printf("Analyzing '%v' with defaults\n", SENTENCE)
	result, err := c.AnalyzeWithDefaults(SENTENCE, "en")
	printResults(result, err, "AnalyzeWithDefaults")

	fmt.Printf("Analyzing '%v' with threshold 0.80\n", SENTENCE)
	var options client.AnalyzerOptions
	options.SetScoreThreshold(0.80)
	result, err = c.AnalyzeWithOptions(SENTENCE, "en", &options)
	printResults(result, err, "AnalyzeWithOptions")
}
