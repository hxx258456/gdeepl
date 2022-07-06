# Copyright 2022 github.com/hxx258456/gdeepl
#
# -------------------------------------------------------------
RELEASE_PLATFORMS = linux-amd64 darwin-amd64 windows-amd64
RELEASE_PKGS = gdeepl
BASE_VERSION = 1.0.0
IS_RELEASE = true
GO_SOURCE := .

release-all: $(patsubst %,release/%, $(RELEASE_PLATFORMS))
release/windows-amd64: GOOS=windows
release/windows-amd64: $(patsubst %,release/windows-amd64/bin/%, $(RELEASE_PKGS))

release/darwin-amd64: GOOS=darwin
release/darwin-amd64: $(patsubst %,release/darwin-amd64/bin/%, $(RELEASE_PKGS))

release/linux-amd64: GOOS=linux
release/linux-amd64: $(patsubst %,release/linux-amd64/bin/%, $(RELEASE_PKGS))

release/%-amd64: GOARCH=amd64

release/linux-%: GOOS=linux

release/%/bin/gdeepl: GO_TAGS+= caclient
release/%/bin/gdeepl: $(GO_SOURCE)
	@echo "Building $@ for $(GOOS)-$(GOARCH)"
	mkdir -p $(@D)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(abspath $@) $(GO_SOURCE)


%-release-clean:
	$(eval TARGET = ${patsubst %-release-clean,%,${@}})
	-@rm -rf release/$(TARGET)
clean: $(patsubst %,%-release-clean, $(RELEASE_PLATFORMS))