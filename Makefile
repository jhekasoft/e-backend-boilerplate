#!/usr/bin/env make
# e-backend-boilerplate

BUILD_TIME=$(shell date -Iseconds)
VERSION=$(shell cat ./VERSION)
LDFLAGS=-s -w -X 'e-backend-boilerplate/internal.BuildTime=$(BUILD_TIME)' -X 'e-backend-boilerplate/internal.Version=$(VERSION)'
TAGS=all
DEV_TAGS=all dev

all: clean build doc data

all-no-workspace: clean build-no-workspace doc data

build:
	go build -ldflags "$(LDFLAGS)" -tags="$(TAGS)" -o ./build/e-backend-boilerplate

build-no-workspace:
	GOWORK=off go build -ldflags "$(LDFLAGS)" -tags="$(TAGS)" -o ./build/e-backend-boilerplate

doc:
	# Update REST doc's version
	sed -i "s/\(version:\) .*/\1 $(VERSION)/" ./modules/doc/data/public/restapi/openapi/openapi.yml

	# Update MQTT doc's version
	sed -i "s/\(version:\) .*/\1 $(VERSION)/" ./modules/doc/data/public/mqttapi/asyncapi/asyncapi.yml

data:
	# Config example
	cp ./config.yaml.example ./build/config.yaml.example

	# Module CV
	mkdir -p ./build/modules/cv/data
	cp -r ./modules/cv/data/* ./build/modules/cv/data
	mkdir -p ./build/modules/cv/templates
	cp -r ./modules/cv/templates/* ./build/modules/cv/templates

	# Module Doc
	mkdir -p ./build/modules/doc/data
	cp -r ./modules/doc/data/* ./build/modules/doc/data

clean:
	rm -rf ./build

run:
	go run -ldflags "$(LDFLAGS)" -tags="$(DEV_TAGS)" main.go serve

run-no-workspace:
	GOWORK=off go run -ldflags "$(LDFLAGS)" -tags="$(DEV_TAGS)" main.go serve

test:
	go test ./...

.PHONY: all doc run
