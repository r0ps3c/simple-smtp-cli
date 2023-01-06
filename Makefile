BINARIES=simple-smtp-cli

all: $(BINARIES)

% :: %.go
	CGO_ENABLED=0 go build -tags osusergo,netgo $<
	strip $@

clean:
	go clean

.PHONY: build all