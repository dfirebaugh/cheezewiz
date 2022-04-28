PROJECTNAME=cheezewiz

all: client server windows-client-cross-compile windows-server-windows-cross-compile wasm web

run:
	go run ./cmd/client

test:
	go test ./...

client:
	go build -o .dist/$(PROJECTNAME) ./cmd/client

server:
	go build -o .dist/$(PROJECTNAME)-server ./cmd/server

windows-client: ## build from windows
	go build -o .dist/$(PROJECTNAME).exe ./cmd/client

windows-client-cross-compile: # cross-compile to windows exe
	GOOS=windows go build -o .dist/$(PROJECTNAME).exe ./cmd/client

windows-server-windows-cross-compile: # cross-compile to windows exe
	GOOS=windows go build -o .dist/$(PROJECTNAME)-server.exe ./cmd/server

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
