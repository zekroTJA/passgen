PACKAGE    = github.com/zekroTJA/passgen
GOPATH     = $(CURDIR)/.gopath
BASE       = $(GOPATH)/src/$(PACKAGE)
BINARY     = passgen
INSTALLDIR = /usr/bin

ifeq ($(GO),)
	GO = go
endif

ifeq ($(OS),Windows_NT)
	WINDOWS = .exe
endif

.PHONY: install clean get move

$(BINARY)$(WINDOWS): $(BASE) get move clean

$(BASE):
	@echo [ INFO ] creating temporary gopath '$@'...
	@mkdir -p $@
	@cp $(CURDIR)/* $@/

get:
	@echo [ INFO ] getting packages and building binaries...
	$(GO) get -v -t $(BASE)/.

move:
	@mv $(GOPATH)/bin/* $(CURDIR)

_install: $(BASE) get
	@echo [ INFO ] installing binaries to '$(INSTALLDIR)/$(BINARY)$(WINDOWS)'...
	@install -m 755 $(GOPATH)/bin/$(BINARY)$(WINDOWS) $(INSTALLDIR)

install: _install clean

clean:
	@echo [ INFO ] cleaning up...
	@rm -r -f $(GOPATH)