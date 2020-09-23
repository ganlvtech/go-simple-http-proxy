# go-simple-http-proxy

## Usage

```bash
./go-simple-http-proxy --help
```

```bash
./go-simple-http-proxy --listen :8080 --remote 127.0.0.1:80
```

```bash
./go-simple-http-proxy --listen :8443 --remote 127.0.0.1:80 --https --cert server.crt --key server.key
```

## Build

```bash
export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GO111MODULE=on
export GOOS=linux
go build -v -ldflags "-s -w" -o .
chmod +x go-simple-udp-logger
```

```batch
SET GOPROXY=https://mirrors.aliyun.com/goproxy/
SET GO111MODULE=on
SET GOOS=windows
go build -v -ldflags "-s -w" .
```
