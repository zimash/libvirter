package libvirter

type Actions interface {
	Connector() // TODO: out to upper level
	Create()
	Delete()
	Start()
	Stop()
	Restart()
	Status()
}
