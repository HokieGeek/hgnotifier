export GOPATH := ${PWD}

install_dir := $(DESTDIR)/$(PREFIX)

all: daemon clients

daemon: src/snotifyd.go src/snotify/snotify.go
	go build -o bin/snotifyd src/snotifyd.go 

clients: src/snotify-capslockd.go src/snotify-message.go
	go build -o bin/snotify-capslockd src/snotify-capslockd.go
	go build -o bin/snotify-message src/snotify-message.go

clean:
	rm -rf bin

install:
	# @echo "TODO: make install target with prefix '$(PREFIX)' in directory '$(DESTDIR)'"
	# @echo "TODO: install_dir = $(install_dir)"
	install -Dm755 bin/snotifyd $(DESTDIR)$(PREFIX)/bin/snotifyd
	install -Dm755 bin/snotify-message $(DESTDIR)$(PREFIX)/bin/snotify-message
	install -Dm755 bin/snotfiy-capslockd $(DESTDIR)$(PREFIX)/bin/snotify-capslockd
	install -Dm644 config/snotify.config $(DESTDIR)$(PREFIX)/etc/snotify.config
	install -Dm644 systemd/snotify.service $(DESTDIR)$(PREFIX)/lib/systemd/system/snotify.service

.PHONY: all daemon clients clean install
