#!/bin/bash

# Step 1: cleanup and download all dependencies
rm -Rf ./main
go mod tidy
go mod vendor

# Step 2: Build the main.go file
go build -o main .

# Step 3: Run the main executable
./main
