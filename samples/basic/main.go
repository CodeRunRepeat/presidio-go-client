package main

import (
	"fmt"

	"github.com/CodeRunRepeat/presidio-go-client/client"
)

func main() {
	var baseUrl string = "https://presidio-analyzer-prod.azurewebsites.net"
	var client = client.NewClient(baseUrl)
	{
		var result, err = client.Health()
		fmt.Printf("Health() returned: %v, %v\n", result, err)
	}
	{
		var result, err = client.AnalyzeWithDefaults("My name is Joe", "en")
		fmt.Printf("AnalyzeWithDefaults() returned: %v, %v\n", result, err)
	}
}
