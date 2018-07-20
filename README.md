# SimpleFileServer

## 编译方法

```sh
GOOS=linux GOARCH=amd64 go build -a -ldflags "-w -s -X main.build=`date '+%m/%d/%Y'` -X main.version=1.0"
```

## 使用方法

```sh
  -v
        version number
  --port int
        listen http port (default 8000)
  --dir string
        listen file path (default .)
```