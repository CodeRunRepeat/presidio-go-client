# Go client for Presidio

[![Go](https://github.com/CodeRunRepeat/presidio-go-client/actions/workflows/go.yml/badge.svg)](https://github.com/CodeRunRepeat/presidio-go-client/actions/workflows/go.yml)

This is an unofficial Go client for [Presidio](https://github.com/microsoft/Presidio), the open source PII detection and anonymization tool. If you are running Presidio in your environment, you can use this client to access its
[REST services](https://microsoft.github.io/presidio/api-docs/api-docs.html) using Go.

Usage examples can be found in the samples folder, and in unit tests.

## Building the client

Open `src` in a dev container and run `make all`. 

## Tools and libraries

* To call Presidio services, the client uses classes generated by [swagger-codegen](https://github.com/swagger-api/swagger-codegen).
* The client uses [staticcheck](https://github.com/dominikh/go-tools) as its linter.
