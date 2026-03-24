# go-wasm-playground

A simple project to learn how Go code can run inside a browser using WebAssembly (WASM) — and actually talk to the page.

---

## What is this?

Normally, browsers only run JavaScript. But with WebAssembly, you can compile Go code into a `.wasm` binary and run it in the browser too.

This project shows the simplest version of that idea: you write a Go function, compile it to WASM, and call it from a webpage. The Go function can read from the page and write back to it — no server needed after the initial load.

---

## How it works

Here's the journey from button click to result:

```
Browser (you click something)
  → JavaScript catches the event
    → calls your Go function (via a JS bridge)
      → Go WASM binary runs in the browser
        → processes the input (e.g. formats JSON)
          → returns the result back to JavaScript /
            or Can manipulate the dom using js bridge too
            → page updates with the output
```

The magic piece in the middle is `wasm_exec.js` — a file that ships with Go. It acts as the bridge that lets JavaScript and your Go binary talk to each other.

---

## Folder structure

```
go-wasm-playground/
├── assets/          # wasm_exec.js and frontend files live here
|── fs.go   # This file embeds the static assets into the binary
└── cmd/
    ├── server/      # tiny Go server to serve the embeds binary
    └── wasm/        # your Go code that gets compiled to .wasm
```

---

## Getting started

**1. Clone the repo**

```bash
git clone git@github.com:SXsid/go-wasm-playground.git && cd go-wasm-playground
```

**2. Copy the WASM bridge file**

Go ships with a file called `wasm_exec.js` that you need to copy into the assets folder:

```bash
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./assets
```

> You only need to do this once. This file is what lets the browser load and talk to your Go binary.

**3. Start the server**

```bash
go run ./cmd/server -port <port>
```

Then open your browser and go to `http://localhost:<port>` (check the terminal for the exact address).

---

## Requirements

- Go 1.25 or later
- A modern browser (Chrome, Firefox, Safari — all support WASM)
