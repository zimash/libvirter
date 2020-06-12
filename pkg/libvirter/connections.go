package libvirter

import (
	"fmt"
	"log"
)

type VMConnection interface {
	Connect() error
	Disconnect() error
	Connected() (bool, error)
	Create(vm VM) error
	Delete(vm VM) error
	Start(vm VM) error
	Stop(vm VM) error
	Restart(vm VM) error
	Status(vm VM) (VMStatus, error)
}

type VMManager struct {
	config      VMManagerConfig
	connections map[VMType]VMConnection
	logger      *log.Logger
}

func NewVMManager(config VMManagerConfig, logger *log.Logger) *VMManager {
	return &VMManager{
		config:      config,
		connections: make(map[VMType]VMConnection),
		logger:      logger,
	}
}

func (vmm *VMManager) GetConnection(vmType VMType) (VMConnection, error) {
	conn, ok := vmm.connections[vmType]
	if !ok {
		switch vmType {
		case VMTypeQEMU:
			conn = NewQEMUConnection()
		default:
			return nil, fmt.Errorf("unsupported vm type %s", vmType)
		}
		vmm.connections[vmType] = conn
	}
	alive, err := conn.Connected()
	if err != nil && err != ErrConnectionIsNotReady {
		return nil, fmt.Errorf("cannot check connection status: %w", err)
	}
	if !alive {
		if err := conn.Connect(); err != nil {
			return nil, fmt.Errorf("cannot connect: %w", err)
		}
	}
	return conn, nil
}
