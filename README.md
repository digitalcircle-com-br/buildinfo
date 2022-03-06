# buildinfo
Build Info for Go Binaries

```makefile
GIT_COMMIT := $(shell git rev-list -1 HEAD)
DT := $(shell date +”%Y.%m.%d.%H%M%S”)
ME := $(shell whoami)
HOST := $(shell hostname)
PRODUCT := "MYAPI"

docker:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X buildinfo.Ver=$(GIT_COMMIT) -X buildinfo.BuildDate=$(DT) -X buildinfo.BuildUser=$(ME) -X buildinfo.BuildHost=$(HOST) -X buildinfo.Product=$(PRODUCT)"  -o deploy/api ./main.go
    cd deploy && \
	docker build -t mymimg .
```