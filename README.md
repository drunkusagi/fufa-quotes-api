# fufa-quotes-api
REST API for random fufufafa quotes.

Thanks to [vanirvan](https://github.com/vanirvan/fufufafa-quotes-fetcher/) for providing the fufufafa quotes JSON file.

## How to build
```bash
# Windows
go build -o fufa-quotes-api.exe -ldflags="-s -w" .\cmd\app\main.go

# Linux
go build -o fufa-quotes-api -ldflags="-s -w" ./cmd/app/main.go
```

## How to run
```bash
# Windows
.\fufa-quotes-api.exe

# Linux
./fufa-quotes-api
```
