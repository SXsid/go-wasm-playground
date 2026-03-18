# go-wasm-playground

### Folder Structure

```
├── assets
└── cmd
    ├── server
    └── wasm
```

## Steps

1. clone the repo

```bash
    git clone git@github.com:SXsid/go-wasm-playground.git && cd go-wasm-palyground

```

2. copy the wasm_exec.js

```bash
 cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./assets


```

2. run

```bash
go run ./cmd/server
```
