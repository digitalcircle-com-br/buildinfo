# buildinfo
Build Info for Go Binaries

```makefile
GIT_COMMIT := $(shell git rev-list -1 HEAD)
DT := $(shell date +”%Y.%m.%d.%H%M%S”)
ME := $(shell whoami)
HOST := $(shell hostname)
PRODUCT := "MYAPI"

docker:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X github.com/digitalcircle-com-br/buildinfo.Ver=$(GIT_COMMIT) -X github.com/digitalcircle-com-br/buildinfo.BuildDate=$(DT) -X github.com/digitalcircle-com-br/buildinfo.BuildUser=$(ME) -X github.com/digitalcircle-com-br/buildinfo.BuildHost=$(HOST) -X github.com/digitalcircle-com-br/buildinfo.Product=$(PRODUCT)"  -o deploy/api ./main.go
    cd deploy && \
	docker build -t mymimg .
```

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "LOCAL",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "buildFlags": "-ldflags '-X github.com/digitalcircle-com-br/buildinfo.Ver=DEV -X github.com/digitalcircle-com-br/buildinfo.BuildDate=Now -X github.com/digitalcircle-com-br/buildinfo.BuildUser=ME -X github.com/digitalcircle-com-br/buildinfo.BuildHost=localhost -X github.com/digitalcircle-com-br/buildinfo.Product=PRODUCT",
            "env": {
                "DSN": "host=localhost user=XXX password=XXX dbname=XXX port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
                "REDIS": "redis://localhost:6379"
            }
        }
    ]
}
```