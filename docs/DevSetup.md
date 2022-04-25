# Dev Setup

1. download and install [golang](https://go.dev/dl/)
2. install `make` (if you're running windows -> [make setup](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows) )
2. Clone this repo
3. navigate to the directory that this repo lives in
4. pull down dependencies

```
go get ./... # pull down dependencies
```



## Run the app

```
make run
```

alternatively:
```
go run ./cmd/client
```