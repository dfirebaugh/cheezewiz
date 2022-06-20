PROJECTNAME=cheezewiz

all: client windows-client-cross-compile wasm web

run:
	go run ./cmd/client

test:
	go test ./...

client:
	go build -o .dist/$(PROJECTNAME) ./cmd/client

windows-client: ## build from windows
	go build -o .dist/$(PROJECTNAME).exe ./cmd/client

windows-client-cross-compile: # cross-compile to windows exe
	GOOS=windows go build -o .dist/$(PROJECTNAME).exe ./cmd/client

wasm:
	GOOS=js GOARCH=wasm go build -o .dist/$(PROJECTNAME).wasm ./cmd/client

.PHONY: web
web:
	mkdir -p .dist \
		&& mkdir -p assets \
		&& cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .dist/ \
		&& cp -R assets .dist/ \
		&& cp -R web/* .dist/ \
		&& mkdir -p .dist/config/entities \
		&& mkdir -p .dist/config/levels \
		&& cp -R config/entities/* .dist/config/entities/ \
		&& cp -R config/levels/* .dist/config/levels/

web-server: ## for testing - note: requires node
	npx es-dev-server --root-dir .dist

clean:
	rm -rf .dist/
