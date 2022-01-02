package main
// import "C"

import (
	"log"
	"github.com/cilium/ebpf"
)

func main() {
	log.Println("loading ....")
	spec, err := ebpf.LoadCollectionSpec("hello.bpf.o")
	if err != nil {
		panic(err)
	}

	// log.Println(spec)
	for k, v := range spec.Programs {
		log.Println(k, v)
	}
}