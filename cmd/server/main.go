package main

import (
	"flag"
	"fmt"
	"net/http"

	root "github.com/SXsid/wasmPlayground"
)

func main() {
	port := flag.Int("port", 8080, "HTTP port address")
	flag.Parse()
	fileHandler := http.FileServerFS(root.AssetsFS)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: fileHandler,
	}

	fmt.Printf("site is up at port %d", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
