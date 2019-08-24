GO_FOLDERS=./evaluation/... ./models/...

pre-commit: go-test go-lint go-doc go-dep-ensure

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

go-bench:
	@go test -benchmem -run='^$$' . -bench '.'

go-lint:
	@golangci-lint run \
	--enable golint --enable gosec --enable interfacer --enable unconvert \
	--enable goimports --enable goconst --enable gocyclo --enable misspell \
	--enable scopelint \
	$(GO_FOLDERS)

go-doc:
	@gfind -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2- | \
		grep -v '^\.' | \
		grep -v '\/\.' | \
		grep -v '^docs$$' | \
		grep -v '^vendor' | \
		xargs -I{} bash -c "godocdown {} > {}/README.md"

go-dep-ensure:
	@dep ensure
