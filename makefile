SHELL := /bin/bash

# ==============================================================================
# Running build within the local computer

build:
	go build -o ./ ./...

# ==============================================================================
# Running tests within the local computer

test:
	go test -count=1 ./...

# ==============================================================================
# Running application within the local computer

run:
	./helloworld-api