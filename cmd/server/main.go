package main

import (
	"fmt"
	"net/http"

	root "github.com/SXsid/wasmPlayground"
)

func main() {
	fileHandler := http.FileServerFS(root.AssetsFS)
	server := http.Server{
		Addr:    ":8080",
		Handler: fileHandler,
	}

	fmt.Println("site is up at port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
