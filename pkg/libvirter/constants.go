package libvirter

const (
	QEMUUri = "qemu:///system"
)

var (
	Flavors = map[string]QEMUVM{
		"tiny":   QEMUVM{CPU: 1, Mem: 512, Disk: 1},
		"small":  QEMUVM{CPU: 1, Mem: 2048, Disk: 20},
		"medium": QEMUVM{CPU: 2, Mem: 4096, Disk: 40},
		"large":  QEMUVM{CPU: 4, Mem: 8192, Disk: 80},
		"xlarge": QEMUVM{CPU: 8, Mem: 16384, Disk: 160},
	}
)
