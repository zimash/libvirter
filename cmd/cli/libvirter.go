package main

import (
	"flag"
	"log"
	"os"

	"github.com/zimash/libvirter/pkg/libvirter"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %s -h\n", os.Args[0])
		os.Exit(0)
	}

	action := flag.String("a", "create", "Action for VM [create, delete, stop, rebot, status]")
	hostname := flag.String("h", "", "VM name")
	cpu := flag.Int("c", 1, "Number of CPU for VM")
	mem := flag.Int("m", 1, "Memory for VM")
	disk := flag.Int("d", 1, "Disk size in Gb")
	flag.Parse()

	vm := libvirter.QEMUVM{
		CPU:  *cpu,
		Name: *hostname,
		Disk: *disk,
		Mem:  *mem,
	}
	vm.Connector()
	switch *action {
	case "create":
	case "delete":
	case "start":
		vm.Start()
	case "stop":
		vm.Stop()
	case "reboot":
		vm.Reboot()
	case "status":
		vm.Status()
	default:
		log.Fatalln("Undefined action:", *action)
	}
}
