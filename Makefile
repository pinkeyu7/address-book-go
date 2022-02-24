# Go parameters
GOCMD:=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
MIGRATE_CONFIG:="./migration/dbconfig.yml"

run:
	$(GORUN) main.go

doc:
	swag init

test:
	$(GOTEST) ./...

migrate-up:
	sql-migrate up -config=$(MIGRATE_CONFIG) -env="localhost"

migrate-down:
	sql-migrate down -config=$(MIGRATE_CONFIG) -env="localhost"
