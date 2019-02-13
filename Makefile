PACKAGE    = github.com/zekroTJA/passgen
GOPATH     = $(CURDIR)/.gopath
BASE       = $(GOPATH)/src/$(PACKAGE)
BINARY     = passgen
INSTALLDIR = /usr/bin

TAG        = $(shell git describe --tags)
COMMIT     = $(shell git rev-parse HEAD)

ifeq ($(GO),)
	GO = go
endif

COMPVER    = $(shell go version | sed -e 's/ /_/g')

ifeq ($(OS),Windows_NT)
	WINDOWS = .exe
endif

.PHONY: install clean get move

_make: $(BINARY)$(WINDOWS) clean

# Creating GOPATH path and copy all files from root path into it
$(BASE):
	@echo [ INFO ] creating temporary gopath '$@'...
	@mkdir -p $@
	@cp $(CURDIR)/* $@/

# Getting dependencies and build binary in current dir
$(BINARY)$(WINDOWS): $(BASE) get
	@echo [ INFO ] building binary '$@'
	$(GO) build -v \
		-ldflags "-X main.ldTag=$(TAG) -X main.ldCommit=$(COMMIT) -X main.ldCompVer=$(COMPVER)" \
		-o $(CURDIR)/$@ $(BASE)/.

get: $(BASE)
	@echo [ INFO ] getting dependencies...
	@cd $(BASE) && $(GO) get -v -t

_install: $(BASE) $(BINARY)$(WINDOWS)
	@echo [ INFO ] installing binaries to '$(INSTALLDIR)/$(BINARY)$(WINDOWS)'...
	@install -m 755 $(CURDIR)/$(BINARY)$(WINDOWS) $(INSTALLDIR)

install: _install clean

clean:
	@echo [ INFO ] cleaning up...
	@rm -r -f $(GOPATH)