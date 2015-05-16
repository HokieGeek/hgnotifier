export GOPATH := ${PWD}

install_dir := $(DESTDIR)/$(PREFIX)

all: snotifyd snotify-message snotify-capslockd 

snotifyd: src/snotifyd.go src/snotify/snotify.go
	go build -o bin/snotifyd src/snotifyd.go 

snotify-message: src/snotify-message.go
	go build -o bin/snotify-message src/snotify-message.go

snotify-capslockd: src/snotify-capslockd.go
	go build -o bin/snotify-capslockd src/snotify-capslockd.go

clean:
	rm -rf bin

install:
	# @echo "TODO: install_dir = $(install_dir)"
	@echo "HERE: ${PWD}"
	ls bin
	install -Dm755 bin/snotifyd $(DESTDIR)$(PREFIX)/bin/snotifyd
	install -Dm755 bin/snotfiy-capslockd $(DESTDIR)$(PREFIX)/bin/snotify-capslockd
	install -Dm755 bin/snotify-message $(DESTDIR)$(PREFIX)/bin/snotify-message
	install -Dm644 config/snotify.config $(DESTDIR)$(PREFIX)/etc/snotify.config
	# install -Dm644 systemd/snotify.service $(DESTDIR)/usr/lib/systemd/system/snotify.service

.PHONY: all clean install
