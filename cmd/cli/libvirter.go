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
	action := flag.String("a", "create", "Action for VM [create, delete, stop, reboot, status]")
	hostname := flag.String("h", "", "VM name")
	cpu := flag.Int("c", 1, "Number of CPU for VM")
	mem := flag.Int("m", 1, "Memory for VM")
	disk := flag.Int("d", 1, "Disk size in Gb")
	vmType := flag.String("t", string(libvirter.VMTypeQEMU), "Type of VM [qemu]")
	flag.Parse()

	config := libvirter.VMManagerConfig{} // TODO: load config file
	manager := libvirter.NewVMManager(config, log.New(os.Stdout, "", log.LstdFlags))

	vm := libvirter.VM{
		CPU:  *cpu,
		Name: *hostname,
		Disk: *disk,
		Mem:  *mem,
	}

	conn, err := manager.GetConnection(libvirter.VMType(*vmType))
	if err != nil {
		log.Fatalln("cannot create connection: ", conn)
	}

	defer func() {
		errConn := conn.Disconnect()
		if errConn != nil {
			log.Println("disconnect error: ", errConn)
		}
	}()

	switch *action {
	case "create":
		err = conn.Create(vm)
	case "delete":
		err = conn.Delete(vm)
	case "start":
		err = conn.Start(vm)
	case "stop":
		err = conn.Stop(vm)
	case "reboot":
		err = conn.Restart(vm)
	case "status":
		var status libvirter.VMStatus
		status, err = conn.Status(vm)
		if err == nil {
			log.Println("status", status)
		}
	default:
		log.Fatalf("undefined action: %s\n", *action)
	}
	if err != nil {
		log.Fatalf("error happened: %v\n", err)
	}
	os.Exit(0)
}
