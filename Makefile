.PHONY: build clean

build:
	go build -o /usr/bin/tammy ./cmd/main.go

clean:
	rm -f $(INSTALL_DIR)/tammy
