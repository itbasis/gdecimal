go-dependencies:
	$(eval GOBIN=$(shell go env GOPATH 2>/dev/null)/bin)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) latest
	#
	go install github.com/editorconfig-checker/editorconfig-checker/v3/cmd/editorconfig-checker@latest
	#
	go install github.com/onsi/ginkgo/v2/ginkgo@latest

go-update: go-dependencies
	go mod tidy && go get -t -v -u ./...

go-generate: go-dependencies
	go generate ./...
	$(MAKE) go-update

go-lint:
	editorconfig-checker
	golangci-lint run

go-test:
	ginkgo -race --cover --coverprofile=.coverage-ginkgo.out --junit-report=junit-report.xml ./...
	go tool cover -func=.coverage-ginkgo.out -o=.coverage.out
	go tool cover -html=.coverage-ginkgo.out -o=coverage.html
	cat .coverage.out

go-all-tests: go-generate go-lint go-test

go-all: go-dependencies go-all-tests
	go mod tidy || :
