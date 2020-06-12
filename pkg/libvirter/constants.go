package libvirter

const (
	VMSizeTiny   VMSize = "tiny"
	VMSizeSmall  VMSize = "small"
	VMSizeMedium VMSize = "medium"
	VMSizeLarge  VMSize = "large"
	VMSizeXLarge VMSize = "xlarge"
)

var (
	Flavors = map[VMSize]VM{
		VMSizeTiny:   {CPU: 1, Mem: 512, Disk: 1},
		VMSizeSmall:  {CPU: 1, Mem: 2048, Disk: 20},
		VMSizeMedium: {CPU: 2, Mem: 4096, Disk: 40},
		VMSizeLarge:  {CPU: 4, Mem: 8192, Disk: 80},
		VMSizeXLarge: {CPU: 8, Mem: 16384, Disk: 160},
	}
)

const (
	VMTypeQEMU VMType = "qemu"
)
