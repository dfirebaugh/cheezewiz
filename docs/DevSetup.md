# Dev Setup

1. download and install [golang](https://go.dev/dl/)
2. install `make` (if you're running windows -> [make setup](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows) )
3. Clone this repo
4. navigate to the directory that this repo lives in
5. pull down dependencies

```
go get ./... # pull down dependencies
```

## Linux

The make file is intended to run on linux.

## Run

```bash
make run
```

### Build

```bash
make
```

### Clean

```bash
make clean
```

### Troubleshoot

If you experience this error

```bash
fatal error: X11/Xlib.h: No such file or directory
#include <X11/Xlib.h>
```

Install the following

```bash
sudo apt install libx11-dev
```

If you experience this error

```bash
# github.com/go-gl/glfw/v3.3/glfw
/usr/bin/ld: cannot find -lXxf86vm
collect2: error: ld returned 1 exit status
make: *** [Makefile:12: client] Error 2
```

Install the following

```bash
sudo apt-get install build-essential libgl1-mesa-dev libglu1-mesa-dev xorg-dev
```

## Windows

On windows, It's probably best to stick to the go commands.
You could setup [WSL](https://docs.microsoft.com/en-us/windows/wsl/install) if you want to test out some of the build commands.

### Run

```powershell
go run ./cmd/client
```

### Build

```powershell
go build -o .dist/cheezewiz.exe ./cmd/client
```

or:

```powershell
make build-win
```

### Clean

```bash
del .dist
```

## Web

You can test out the web build by running:

```
npx es-dev-server --root-dir .dist
```

> Note: this requires that you have [node](https://nodejs.org/en/download/) installed
