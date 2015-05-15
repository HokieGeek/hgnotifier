export GOPATH := ${PWD}

install_dir := $(DESTDIR)/$(PREFIX)

all: daemon clients

daemon: src/snotify-daemon.go src/snotify/snotify.go
	go build -o bin/snotify-daemon src/snotify-daemon.go 

clients: src/capslock-listener.go src/snotify-message.go
	go build -o bin/capslock-listener src/capslock-listener.go
	go build -o bin/snotify-message src/snotify-message.go

clean:
	rm -rf bin

install:
	# @echo "TODO: make install target with prefix '$(PREFIX)' in directory '$(DESTDIR)'"
	# @echo "TODO: install_dir = $(install_dir)"
	install -Dm755 bin/snotify-daemon $(DESTDIR)$(PREFIX)/bin/snotify-daemon
	install -Dm755 bin/snotify-message $(DESTDIR)$(PREFIX)/bin/snotify-message
	install -Dm755 bin/capslock-listener $(DESTDIR)$(PREFIX)/bin/capslock-listener
	install -Dm644 config/snotify.config $(DESTDIR)$(PREFIX)/etc/snotify/snotify.config
	install -Dm644 systemd/snotify.service $(DESTDIR)$(PREFIX)/lib/systemd/system/snotify.service

.PHONY: all daemon clients clean install
