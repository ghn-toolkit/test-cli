.PHONY: clean lint security critic test install build release build-and-push-images delete-tag update-pkg-cache

clean:
	rm -rf ./tmp ./tests

lint:
	$(GOPATH)/bin/golangci-lint run ./...

security:
	$(GOPATH)/bin/gosec -quiet ./...

critic:
	$(GOPATH)/bin/gocritic check -enableAll ./...

test: clean lint security critic
	mkdir ./tests
	go test -coverprofile=./tests/coverage.out ./...
	go tool cover -func=./tests/coverage.out
	rm -rf ./tests

install: test
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cga ./cmd/cga/main.go

build: test
	$(GOPATH)/bin/goreleaser --snapshot --skip-publish --rm-dist

release: test
	git tag -a v$(VERSION) -m "$(VERSION)"
	$(GOPATH)/bin/goreleaser --snapshot --skip-publish --rm-dist

update-pkg-cache:
	curl -i https://proxy.golang.org/github.com/ghn-toolkit/test-cli/@v/v$(VERSION).info

delete-tag:
	git tag --delete v$(VERSION)
