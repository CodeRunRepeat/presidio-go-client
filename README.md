# Go client for Presidio

[![Go](https://github.com/CodeRunRepeat/presidio-go-client/actions/workflows/go.yml/badge.svg)](https://github.com/CodeRunRepeat/presidio-go-client/actions/workflows/go.yml)

This is an unofficial client for [Presidio](https://github.com/microsoft/Presidio), the open source PII detection and anonymization tool. If you are running Presidio in your environment, you can use this client to access its REST services.

Currently, only the analyzer service client is implemented.

## Building the client

Open `src` in a dev container and run `make all`. The client uses [staticcheck](https://github.com/dominikh/go-tools) as its linter.
