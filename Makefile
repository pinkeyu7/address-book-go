# Go parameters
GOCMD:=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

run:
	$(GORUN) main.go

doc:
	swag init

test:
	$(GOTEST) ./...
