#!/bin/bash

# Run all tests
go test ./...

# Build the binary
go build

# Run with test configuration
./hyprmax 