export GOPATH := ${PWD}

install_dir := $(DESTDIR)/$(PREFIX)

all: daemon clients

daemon: src/hgdaemon.go src/hgnotify/hgnotify.go
	go build -o bin/hgdaemon src/hgdaemon.go 

clients: src/capslock-listener.go src/hgmessage.go
	go build -o bin/capslock-listener src/capslock-listener.go
	go build -o bin/hgmessage src/hgmessage.go

clean:
	rm -rf bin

install:
	# @echo "TODO: make install target with prefix '$(PREFIX)' in directory '$(DESTDIR)'"
	# @echo "TODO: install_dir = $(install_dir)"
	install -Dm755 bin/hgdaemon $(DESTDIR)$(PREFIX)/bin/hgdaemon
	install -Dm755 bin/hgmessage $(DESTDIR)$(PREFIX)/bin/hgmessage
	install -Dm755 bin/capslock-listener $(DESTDIR)$(PREFIX)/bin/capslock-listener
	install -Dm644 config/hgnotifier.config $(DESTDIR)$(PREFIX)/etc/hgnotifier/hgnotifier.config
	install -Dm644 systemd/hgnotifier.service $(DESTDIR)$(PREFIX)/lib/systemd/system/hgnotifier.service

.PHONY: all daemon clients clean install
