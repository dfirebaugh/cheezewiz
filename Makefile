PROJECTNAME=cheesewiz

all: build build-windows

run:
	go run ./cmd/client

build:
	go build -o .dist/$(PROJECTNAME) ./cmd/client

build-windows: # cross-compile to windows exe
	GOOS=windows go build -o .dist/$(PROJECTNAME).exe ./cmd/client

test:
	go test ./...
