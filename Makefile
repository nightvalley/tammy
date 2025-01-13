.PHONY: build clean

build:
	@if grep -q 'ID=arch' /etc/os-release; then \
		echo "Detected Arch Linux. Using makepkg."; \
		makepkg -si --clean --noconfirm; \
		rm -rf tammy*; \
	else \
		echo "Detected non-Arch Linux or macOS. Using go build."; \
		go build -o /usr/bin/tammy ./cmd/tammy/main.go; \
	fi

clean:
	rm -f /usr/bin/tammy
