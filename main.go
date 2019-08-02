package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	ip := flag.String("i", "0.0.0.0", "the listen ip")
	port := flag.String("p", "8080", "the listen port")
	flag.Parse()

	WorkDir, _ := filepath.Abs(".")
	log.Println(WorkDir)

	log.Printf("String server listening: %s:%s\n", *ip, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *ip, *port), http.FileServer(http.Dir(WorkDir))))
}
