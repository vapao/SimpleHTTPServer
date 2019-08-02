package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func main() {
	ip := flag.String("i", "0.0.0.0", "the listen ip")
	port := flag.String("p", "8080", "the listen port")
	flag.Parse()

	file, _ := exec.LookPath(os.Args[0])
	WorkDir := path.Dir(file)

	log.Printf("String server listening: %s:%s\n", *ip, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *ip, *port), http.FileServer(http.Dir(WorkDir))))
}
