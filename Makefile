.DEFAULT_GOAL := help

NAME := hcl-util

VERSION := $(shell git describe --tags --abbrev=0)
VERSION_LONG := $(shell git describe --tags)
VAR_VERSION := github.com/tmtk75/hcl-util/cmd.Version

LDFLAGS := -ldflags "-X $(VAR_VERSION)=$(VERSION) \
	-X $(VAR_VERSION)Long=$(VERSION_LONG)"

SRCS := $(shell find . -type f -name '*.go')

.PHONY: build
build: hcl-util ## Build here

hcl-util: $(SRCS)
	go build $(LDFLAGS) -o $(NAME) main.go


.PHONY: install
install:  ## Install in GOPATH
	go install $(LDFLAGS) main.go

.PHONY: clean
clean:  ## Clean
	rm -f $(NAME)

distclean: clean
	rm -rf build vendor


## Release targets
.PHONY: release
release: upload-archives

.PHONY: upload-archives
upload-archives: archive
	echo ghr -u tmtk75 $(VERSION) ./build/*.zip

.PHONY: archive
archive: release-build
	for n in linux windows darwin; do \
	  (cd build; zip $(NAME)_$${n}_amd64.zip $(NAME)_$${n}_amd64) \
	done

.PHONY: release-build
release-build:
	for n in linux windows darwin; do \
	  (GOARCH=amd64 GOOS=$${n} go build -o build/$(NAME)_$${n}_amd64 main.go) \
	done

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
