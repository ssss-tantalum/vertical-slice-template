# Golang Vertical Slice Architecture  Template

A vertical slice architecture Go backend template built with Echo, Huma, MySQL(MariaDB), and Uber FX.

## Tools

- Echo (web framework)
- Huma (OpenAPI3 Documentation)
- FX (DI)
- gomock (Mock)
- Viper (configs)

## How to start application

1. clone repo

`git clone https://github.com/ssss-tantalum/vertical-slice-template.git`

2. copy env variables to .env

`cd vertical-slice-template`

`cp .devcontainer/.env.example .devcontainer/.env`

2. build a dev container (in VSCode)

`code .`

ctrl+P -> "Dev Containers: Open Folder in Container..."

3. start server

`make dev`

## Server

### Options

```bash
$ go run cmd/cli/main.go --help
Usage:
  main [flags]

Flags:
      --env string    (default "dev")
  -h, --help         help for main
      --port int      (default 8000)
```

### Base URI

- API
    - `http://localhost:8000/api`
- Doc
    - `http://localhost:8000/docs`

## Folder Stracture

```
├── cmd
│   └── cli
│       └── main.go
├── configs
│   └── dev.yaml
├── go.mod
├── go.sum
├── internal
│   ├── features
│   │   └── task
│   │       ├── handler
│   │       │   ├── create.go
│   │       │   ├── delete.go
│   │       │   ├── handler.go
│   │       │   ├── list.go
│   │       │   ├── show.go
│   │       │   └── update.go
│   │       ├── infrastracture
│   │       │   ├── create.go
│   │       │   ├── delete.go
│   │       │   ├── find.go
│   │       │   ├── repository.go
│   │       │   └── update.go
│   │       ├── mock
│   │       │   └── task_mock.go
│   │       ├── service
│   │       │   ├── create.go
│   │       │   ├── create_test.go
│   │       │   ├── delete.go
│   │       │   ├── delete_test.go
│   │       │   ├── find.go
│   │       │   ├── find_test.go
│   │       │   ├── service.go
│   │       │   ├── update.go
│   │       │   └── update_test.go
│   │       └── task.go
│   └── shared
│       ├── api
│       │   └── api.go
│       ├── config
│       │   └── config.go
│       ├── database
│       │   └── database.go
│       ├── router
│       │   └── router.go
│       └── server
│           └── server.go
├── Makefile
└── README.md
```

