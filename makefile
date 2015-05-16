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
	@echo "Installing daemons"
	install -Dm755 bin/snotifyd $(DESTDIR)$(PREFIX)/bin/snotifyd
	install -Dm755 bin/snotify-capslockd $(DESTDIR)$(PREFIX)/bin/snotify-capslockd
	install -Dm755 bin/snotify-message $(DESTDIR)$(PREFIX)/bin/snotify-message
	@echo "Installing configs"
	install -Dm644 config/snotify.config $(DESTDIR)$(PREFIX)/etc/snotify.config
	# install -Dm644 systemd/snotify.service $(DESTDIR)/usr/lib/systemd/system/snotify.service
	@echo "Installing notifiers"
	install -Dm755 notifiers/bluetooth-osd.sh $(DESTDIR)$(PREFIX)/share/snotify/notifiers/bluetooth-osd.sh
	install -Dm644 notifiers/bt.xbm $(DESTDIR)$(PREFIX)/share/snotify/notifiers/bt.xbm
	install -Dm755 notifiers/capslock-osd.sh $(DESTDIR)$(PREFIX)/share/snotify/notifiers/capslock-osd.sh
	install -Dm755 notifiers/lowbattery-osd.sh $(DESTDIR)$(PREFIX)/share/snotify/notifiers/lowbattery-osd.sh

.PHONY: all clean install
