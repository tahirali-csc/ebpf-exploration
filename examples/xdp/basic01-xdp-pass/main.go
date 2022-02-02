package main

import "C"
import (
	"bytes"
	_ "embed"
	"fmt"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/rlimit"
	"github.com/dropbox/goebpf"
	"github.com/vishvananda/netlink"
)

//go:embed xdp_prog_kern.o
var xdpPassKen []byte

func main111() {
	// In order to be simple this examples does not handle errors
	bpf := goebpf.NewDefaultEbpfSystem()
	// Read clang compiled binary
	err := bpf.LoadElf("xdp_prog_kern.o")
	if err != nil {
		panic(err)
	}

	log.Println(bpf)
	for k, v := range bpf.GetPrograms() {
		log.Println(k, v)
	}
	// Load XDP program into kernel (name matches function name in C)
	// xdp := bpf.GetProgramByName("xdp_drop_func")
	// // log.Println(xdp)
	// xdp.Load()
	// // Attach to interface
	// xdp.Attach("veth-basic02")

	// xdp = bpf.GetProgramByName("xdp_drop_func")
	// xdp.Load()
	// // Attach to interface
	// xdp.Attach("veth-basic02")

	// defer xdp.Detach()
	// // Work with maps
	// test := bpf.GetMapByName("test")
	// value, _ := test.LookupInt(0)
	// fmt.Printf("Value at index 0 of map 'test': %d\n")
}

type XdpAttachMode int

const (
	// XdpAttachModeNone stands for "best effort" - kernel automatically
	// selects best mode (would try Drv first, then fallback to Generic).
	// NOTE: Kernel will not fallback to Generic XDP if NIC driver failed
	//       to install XDP program.
	XdpAttachModeNone XdpAttachMode = 0
	// XdpAttachModeSkb is "generic", kernel mode, less performant comparing to native,
	// but does not requires driver support.
	XdpAttachModeSkb XdpAttachMode = (1 << 1)
	// XdpAttachModeDrv is native, driver mode (support from driver side required)
	XdpAttachModeDrv XdpAttachMode = (1 << 2)
	// XdpAttachModeHw suitable for NICs with hardware XDP support
	XdpAttachModeHw XdpAttachMode = (1 << 3)
)

func main() {
	// Allow the current process to lock memory for eBPF resources.
	// if err := rlimit.RemoveMemlock(); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(len(xdpPassKen))
	spec, err := loadXdp1()
	if err != nil {
		panic(err)
	}

	objs := &xdpObjects{}
	spec.LoadAndAssign(objs, nil)

	defer objs.xdpPrograms.Close()
	defer objs.xdpMaps.Close()

	// log.Println(objs)

	link, err := netlink.LinkByName("veth-basic02")
	if err != nil {
		panic(err)
	}

	log.Println("XdpAbortFunc=", objs.XdpAbortFunc.FD())
	log.Println("XdpDropFunc=", objs.XdpDropFunc.FD())
	log.Println("XdpPassFunc=", objs.XdpPassFunc.FD())

	// err = netlink.LinkSetXdpFdWithFlags(link, 3, int(XdpAttachModeDrv|XdpAttachModeNone))
	// if err != nil {
	// 	panic(err)
	// }
	err = netlink.LinkSetXdpFdWithFlags(link, 4, int(XdpAttachModeDrv|XdpAttachModeNone))
	if err != nil {
		panic(err)
	}
	err = netlink.LinkSetXdpFdWithFlags(link, 5, int(XdpAttachModeDrv|XdpAttachModeNone))
	if err != nil {
		panic(err)
	}

}

func loadXdp1() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(xdpPassKen)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load xdp: %w", err)
	}

	return spec, err
}

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc $BPF_CLANG -cflags $BPF_CFLAGS -target bpf xdp xdp_pass_kern.c -- -I./
func main1111() {

	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	// log.Println(len(_XdpBytes))

	objs := &xdpObjects{}
	err := loadXdpObjects(objs, nil)
	if err != nil {
		panic(err)
	}
	// log.Println(objs.XdpAbortFunc)

	link, err := netlink.LinkByName("veth-basic02")
	if err != nil {
		panic(err)
	}
	// log.Println(objs.XdpAbortFunc.FD())
	err = netlink.LinkSetXdpFdWithFlags(link, objs.XdpAbortFunc.FD(), 2)
	if err != nil {
		panic(err)
	}
	err = netlink.LinkSetXdpFdWithFlags(link, objs.XdpDropFunc.FD(), 2)
	if err != nil {
		panic(err)
	}
	err = netlink.LinkSetXdpFdWithFlags(link, objs.XdpPassFunc.FD(), 2)
	if err != nil {
		panic(err)
	}

}
