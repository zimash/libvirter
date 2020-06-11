package libvirter

type VMType string

type VMSize string

type VM struct {
	Name string
	Disk int
	CPU  int
	Mem  int
}

type VMStatus struct {
	Active bool
}
