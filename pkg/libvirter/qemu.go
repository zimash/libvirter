package libvirter

import (
	"log"

	libvirt "github.com/libvirt/libvirt-go"
)

type QEMUVM struct {
	Name string
	Disk int
	CPU  int
	Mem  int
	Conn *libvirt.Connect
}

func (q *QEMUVM) Connector() {
	conn, err := libvirt.NewConnect(QEMUUri)
	if err != nil {
		log.Fatal(err)
	}
	q.Conn = conn

	/*
		defer func() {
			ret, err := q.Conn.Close()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(ret)
		}()
	*/
}

func (q *QEMUVM) Create() {
	if q.Conn == nil {
		log.Fatalln("libvirt.Connect is nil")
	}

}

func (q *QEMUVM) Delete() {
	if q.Conn == nil {
		log.Fatalln("libvirt.Connect is nil")
	}
}

func (q *QEMUVM) Start() {
	if q.Conn == nil {
		log.Fatalln("libvirt.Connect is nil")
	}

	dom, err := q.Conn.LookupDomainByName(q.Name)
	if err != nil {
		log.Fatalln(err)
	}
	err = dom.Create()
	if err != nil {
		log.Fatalln(err)
	}
}

func (q *QEMUVM) Stop() {
	if q.Conn == nil {
		log.Fatalln("libvirt.Connect is nil")
	}

	dom, err := q.Conn.LookupDomainByName(q.Name)
	if err != nil {
		log.Fatalln(err)
	}
	err = dom.Shutdown()
	if err != nil {
		log.Fatalln(err)
	}
}

func (q *QEMUVM) Reboot() {
	if q.Conn == nil {
		log.Fatalln("libvirt.Connect is nil")
	}
	dom, err := q.Conn.LookupDomainByName(q.Name)
	if err != nil {
		log.Fatalln(err)
	}
	err = dom.Reboot(libvirt.DOMAIN_REBOOT_ACPI_POWER_BTN)
	if err != nil {
		log.Fatalln(err)
	}
}

func (q *QEMUVM) Status() {
	if q.Conn == nil {
		log.Fatalln("libvirt.Connect is nil")
	}

	dom, err := q.Conn.LookupDomainByName(q.Name)
	if err != nil {
		log.Fatalln(err)
	}
	active, err := dom.IsActive()
	if err != nil {
		log.Fatalln(err)
	}
	if active {
		log.Println("VM", q.Name, "is started")
	} else {
		log.Println("VM", q.Name, "is not started")
	}
}
