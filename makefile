export GOPATH := ${PWD}

all: bin/snotifyd bin/snotify-message bin/snotify-capslockd 

bin/snotifyd: src/snotifyd.go src/snotify/snotify.go
	go build -o $@ src/snotifyd.go 

bin/snotify-message: src/snotify-message.go src/snotify/snotify.go
	go build -o $@ src/snotify-message.go

bin/snotify-capslockd: src/snotify-capslockd.go src/snotify/snotify.go
	go build -o $@ src/snotify-capslockd.go

clean:
	rm -rf bin

install:
	@echo "Installing daemons"
	install -Dm755 bin/snotifyd $(DESTDIR)$(PREFIX)/bin/snotifyd
	install -Dm755 scripts/snotifyctl $(DESTDIR)$(PREFIX)/bin/snotifyctl
	install -Dm755 bin/snotify-message $(DESTDIR)$(PREFIX)/bin/snotify-message
	install -Dm755 bin/snotify-capslockd $(DESTDIR)$(PREFIX)/bin/snotify-capslockd
	@echo "Installing config"
	install -Dm644 etc/snotify.config $(DESTDIR)$(PREFIX)/etc/snotify.config
	@echo "Installing notifiers"
	install -Dm755 notifiers/bluetooth-osd.sh $(DESTDIR)$(PREFIX)/share/snotify/notifiers/bluetooth-osd.sh
	install -Dm644 notifiers/bt.xbm $(DESTDIR)$(PREFIX)/share/snotify/notifiers/bt.xbm
	install -Dm755 notifiers/capslock-osd.sh $(DESTDIR)$(PREFIX)/share/snotify/notifiers/capslock-osd.sh
	install -Dm755 notifiers/lowbattery-osd.sh $(DESTDIR)$(PREFIX)/share/snotify/notifiers/lowbattery-osd.sh
	@echo "Installing triggers"
	install -Dm755 triggers/bluetooth-state.sh $(DESTDIR)$(PREFIX)/share/snotify/triggers/bluetooth-state.sh
	install -Dm755 triggers/low-battery.sh $(DESTDIR)$(PREFIX)/share/snotify/triggers/low-battery.sh

uninstall:
	@echo "Uninstalling daemons"
	rm -rf $(DESTDIR)$(PREFIX)/bin/snotifyd
	rm -rf $(DESTDIR)$(PREFIX)/bin/snotifyctl
	rm -rf $(DESTDIR)$(PREFIX)/bin/snotify-message
	rm -rf $(DESTDIR)$(PREFIX)/bin/snotify-capslockd
	@echo "Uninstalling configs"
	rm -rf $(DESTDIR)$(PREFIX)/etc/snotify.config
	@echo "Uninstalling notifiers"
	rm -rf $(DESTDIR)$(PREFIX)/share/snotify/notifiers/bluetooth-osd.sh
	rm -rf $(DESTDIR)$(PREFIX)/share/snotify/notifiers/bt.xbm
	rm -rf $(DESTDIR)$(PREFIX)/share/snotify/notifiers/capslock-osd.sh
	rm -rf $(DESTDIR)$(PREFIX)/share/snotify/notifiers/lowbattery-osd.sh
	@echo "Uninstalling triggers"
	rm -rf $(DESTDIR)$(PREFIX)/share/snotify/triggers/bluetooth-state.sh
	rm -rf $(DESTDIR)$(PREFIX)/share/snotify/triggers/low-battery.sh

.PHONY: all clean install uninstall
