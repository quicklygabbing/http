GOPATH ?= $(HOME)/go
LINTER=$(GOPATH)/bin/golangci-lint

.PHONY: deps githooks lint test test_stop test_down dev_up dev_stop dev_down

$(LINTER):
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1

deps: githooks
	GO111MODULE=on go mod graph
	GO111MODULE=on go mod vendor

lint: $(LINTER)
	GO111MODULE=on $(GOPATH)/bin/golangci-lint run

http: *.go cmd/*.go
	CGO_ENABLED=0 go build -o http .

test: test_down
	docker network create quicklygabbing-test || true
	docker-compose -f deployments/docker-compose-test.yaml build --no-cache quicklygabbing-http-test-api
	docker-compose -f deployments/docker-compose-test.yaml run quicklygabbing-http-test

test_stop:
	docker-compose -f ./deployments/docker-compose-test.yaml stop

test_down:
	docker-compose -f ./deployments/docker-compose-test.yaml down --remove-orphans
	docker network rm quicklygabbing-test || true

dev_up: dev_down
	docker network create quicklygabbing-dev || true
	docker-compose -f deployments/docker-compose.yaml up --force-recreate --build -d

dev_stop:
	docker-compose -f ./deployments/docker-compose.yaml stop

dev_down:
	docker-compose -f ./deployments/docker-compose.yaml down --remove-orphans
	docker network remove quicklygabbing-dev || true

.git/hooks:
	mkdir -p .git/hooks

.git/hooks/pre-push: githooks/pre-push
	cp githooks/pre-push .git/hooks/pre-push

.git/hooks/pre-commit: githooks/pre-commit
	cp githooks/pre-commit .git/hooks/pre-commit

githooks: .git/hooks .git/hooks/pre-push .git/hooks/pre-commit $(LINTER)
	chmod +x .git/hooks/*
