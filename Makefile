.PHONY: build test run clean

build:
	go build

test:
	go test ./...

run: build
	./hyprmax

clean:
	rm -f hyprmax
	go clean

install: build
	cp hyprmax $(GOPATH)/bin/ 