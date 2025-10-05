module example.com/samples/analyzer-advanced

go 1.23.0

toolchain go1.24.1

require github.com/CodeRunRepeat/presidio-go-client v0.0.1

require (
	github.com/antihax/optional v1.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	golang.org/x/oauth2 v0.27.0 // indirect
)

replace github.com/CodeRunRepeat/presidio-go-client => ../../src
