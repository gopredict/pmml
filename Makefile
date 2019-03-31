GO_FOLDERS=.

pre-commit: go-test go-lint go-dep-ensure

commit:
	@git cz

release:
	@semantic-release

utilities: setup
	@npm install -g commitizen cz-conventional-changelog @commitlint/prompt-cli @commitlint/config-conventional
	@npm install -g semantic-release @semantic-release/commit-analyzer @semantic-release/release-notes-generator @semantic-release/changelog @semantic-release/git

setup:
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

go-test:
	@go test $(GO_FOLDERS)

go-lint:
	@golangci-lint run \
	--enable golint --enable gosec --enable interfacer --enable unconvert \
	--enable goimports --enable goconst --enable gocyclo --enable misspell \
	--enable scopelint \
	$(GO_FOLDERS)

go-dep-ensure:
	@dep ensure
