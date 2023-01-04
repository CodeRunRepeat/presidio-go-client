package main

import (
	"fmt"

	"github.com/CodeRunRepeat/presidio-go-client/client"
)

func main() {
	var baseUrl string = "https://presidio-analyzer-prod.azurewebsites.net"

	basicSample(baseUrl)

	basicAuthSample(baseUrl)
}

func basicSample(baseUrl string) {
	var client = client.NewClient(baseUrl, nil)
	{
		var result, err = client.Health()
		fmt.Printf("Health() returned: %v, %v\n", result, err)
	}
	{
		var result, err = client.AnalyzeWithDefaults("My name is Joe", "en")
		fmt.Printf("AnalyzeWithDefaults() returned: %v, %v\n", result, err)
	}
}

func basicAuthSample(baseUrl string) {
	var client = client.NewClient(baseUrl, client.BasicAuth{UserName: "test", Password: "pass@word1"}) /* imaginary credentials*/

	{
		var result, err = client.Health()
		fmt.Printf("Health() returned: %v, %v\n", result, err)
	}
}
