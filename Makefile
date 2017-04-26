WANT_VERSION = $(shell grep '^v[0-9]' CHANGELOG.md | head -n1 | cut -d' ' -f1)

# Minor versions of Go for which the lint check should be run.
LINTABLE_MINOR_VERSIONS := 7

# Paths besides auto-detected generated files that should be excluded from
# lint results. If generated code uses '//line' directives with a different
# filename, that file won't be auto-detected.
LINT_EXCLUDES_EXTRAS = \
	idl/internal/lex.rl \
	idl/internal/thrift.y \
	idl/internal/yaccpar \
	compile/primitive.go

ERRCHECK_FLAGS ?= -ignoretests

##############################################################################
export GO15VENDOREXPERIMENT=1

GO_VERSION := $(shell go version | cut -d' ' -f3)   # e.g.: go1.6.2
GO_MINOR_VERSION := $(word 2, $(subst ., , $(GO_VERSION)))

ifneq ($(filter $(LINTABLE_MINOR_VERSIONS),$(GO_MINOR_VERSION)),)
SHOULD_LINT := true
endif

PACKAGES := $(shell glide novendor)

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

# Files whose first line contains "Code generated by" are generated.
GENERATED_GO_FILES := $(shell \
	find $(GO_FILES) \
	-exec sh -c 'head -n1 {} | grep "Code generated by\|Automatically generated by " >/dev/null' \; \
	-print)

LINT_EXCLUDES := $(GENERATED_GO_FILES) $(LINT_EXCLUDES_EXTRAS)

# Pipe lint output into this to filter out ignored files.
FILTER_LINT := grep -v $(patsubst %,-e %, $(LINT_EXCLUDES)) -e "vendor/"

# Disable printf-like invocation checking due to testify.assert.Error()
VET_RULES := -printf=false

BUILD_FLAGS ?=

.PHONY: build
build:
	go build -i $(BUILD_FLAGS)

.PHONY: generate
generate:
	go build -i -tags=thriftrw.disableVersionCheck
	PATH=$$(pwd):$$PATH go generate $$(glide nv)
	make -C ./gen/testdata
	./scripts/updateLicenses.sh

.PHONY: lint
lint:
ifdef SHOULD_LINT
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@gofmt -e -s -l $(GO_FILES) | $(FILTER_LINT) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

	$(eval VET_LOG := $(shell mktemp -t govet.XXXXX))
	@$(foreach pkg, $(PACKAGES), \
		go tool vet $(VET_RULES) $(pkg) 2>&1 | \
			grep -v '^exit status' | \
			$(FILTER_LINT) | \
			grep -v 'something"` not compatible with reflect.StructTag.Get:' \
		> $(VET_LOG) || true; \
	)
	@[ ! -s "$(VET_LOG)" ] || (echo "govet failed:" | cat - $(VET_LOG) && false)

	$(eval LINT_LOG := $(shell mktemp -t golint.XXXXX))
	@cat /dev/null > $(LINT_LOG)
	@$(foreach pkg, $(PACKAGES), golint $(pkg) | $(FILTER_LINT) >> $(LINT_LOG) || true;)
	@[ ! -s "$(LINT_LOG)" ] || (echo "golint failed:" | cat - $(LINT_LOG) && false)

	$(eval ERRCHECK_LOG := $(shell mktemp -t errcheck.XXXXX))
	@cat /dev/null > $(ERRCHECK_LOG)
	@$(foreach pkg, $(PACKAGES), errcheck $(ERRCHECK_FLAGS) $(pkg) | $(FILTER_LINT) >> $(ERRCHECK_LOG) || true;)
	@[ ! -s "$(ERRCHECK_LOG)" ] || (echo "errcheck failed:" | cat - $(ERRCHECK_LOG) && false)
else
	@echo "Skipping linters for $(GO_VERSION)"
endif

.PHONY: verifyVersion
verifyVersion: build
	@if [ "$$(./thriftrw --version)" != "thriftrw $(WANT_VERSION)" ]; then \
		echo "Version number in version.go does not match CHANGELOG.md"; \
		echo "Want: thriftrw $(WANT_VERSION)"; \
		echo " Got: $$(./thriftrw --version)"; \
		exit 1; \
	else \
		echo "thriftrw $(WANT_VERSION)"; \
	fi

.PHONY: test
test: build verifyVersion
	go test -race $(PACKAGES)

.PHONY: cover
cover:
	./scripts/cover.sh $(shell go list $(PACKAGES))
	go tool cover -html=cover.out -o cover.html

.PHONY: clean
clean:
	go clean
	rm -rf cover/cover*.out cover.html cover.out

.PHONY: install
install:
	glide --version || go get github.com/Masterminds/glide
	glide install

##############################################################################
# CI

.PHONY: install_ci
install_ci: install
ifdef SHOULD_LINT
	go get -u -f github.com/golang/lint/golint
	go get -u -f github.com/kisielk/errcheck
endif
	go get -u github.com/wadey/gocovmerge
	go get -u github.com/mattn/goveralls
	go get -u golang.org/x/tools/cmd/cover
	go install .

.PHONY: build_ci
build_ci: build

.PHONY: lint_ci
lint_ci: lint

.PHONY: test_ci
test_ci: build_ci verifyVersion
	./scripts/cover.sh $(shell go list $(PACKAGES))
