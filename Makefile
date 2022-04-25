PROJECTNAME=cheezewiz

all: build build-windows-cross wasm web

run:
	go run ./cmd/client

build:
	go build -o .dist/$(PROJECTNAME) ./cmd/client

build-win: ## build from windows
	go build -o .dist/$(PROJECTNAME).exe ./cmd/client

build-windows-cross: # cross-compile to windows exe
	GOOS=windows go build -o .dist/$(PROJECTNAME).exe ./cmd/client

test:
	go test ./...

wasm:
	GOOS=js GOARCH=wasm go build -o .dist/$(PROJECTNAME).wasm ./cmd/client

.PHONY: web
web:
	mkdir -p .dist \
		&& mkdir -p assets \
		&& cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .dist/ \
		&& cp -R assets .dist/ \
		&& cp web/index.html .dist/ \
		&& cp web/main.html .dist/

web-server: ## for testing - note: requires node
	npx es-dev-server --root-dir .dist

clean:
	rm -rf .dist/
