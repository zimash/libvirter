package libvirter

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

const (
	QEMUUri = "qemu:///system"
)

type QEMUConnection struct {
	conn *libvirt.Connect
}

func NewQEMUConnection() *QEMUConnection {
	return &QEMUConnection{}
}

func (c *QEMUConnection) Connect() error {
	var err error
	c.conn, err = libvirt.NewConnect(QEMUUri)
	if err != nil {
		return err
	}
	return nil
}

func (c *QEMUConnection) Disconnect() error {
	if c.conn == nil {
		return ErrConnectionIsNotReady
	}
	_, err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *QEMUConnection) Connected() (bool, error) {
	if c.conn == nil {
		return false, ErrConnectionIsNotReady
	}
	return c.conn.IsAlive()
}

func (c *QEMUConnection) Create(vm VM) error {
	if c.conn == nil {
		return ErrConnectionIsNotReady
	}
	panic("implement me")
}

func (c *QEMUConnection) Delete(vm VM) error {
	if c.conn == nil {
		return ErrConnectionIsNotReady
	}
	panic("implement me")
}

func (c *QEMUConnection) Start(vm VM) error {
	if c.conn == nil {
		return ErrConnectionIsNotReady
	}
	dom, err := c.conn.LookupDomainByName(vm.Name)
	if err != nil {
		return fmt.Errorf("cannot find vm by name: %w", err)
	}
	err = dom.Create()
	if err != nil {
		return fmt.Errorf("cannot create vm: %w", err)
	}
	return nil
}

func (c *QEMUConnection) Stop(vm VM) error {
	if c.conn == nil {
		return ErrConnectionIsNotReady
	}
	dom, err := c.conn.LookupDomainByName(vm.Name)
	if err != nil {
		return fmt.Errorf("cannot find vm by name: %w", err)
	}
	err = dom.Shutdown()
	if err != nil {
		return fmt.Errorf("cannot shutdown vm: %w", err)
	}
	return nil
}

func (c *QEMUConnection) Restart(vm VM) error {
	if c.conn == nil {
		return ErrConnectionIsNotReady
	}
	dom, err := c.conn.LookupDomainByName(vm.Name)
	if err != nil {
		return fmt.Errorf("cannot find vm by name: %w", err)
	}
	err = dom.Reboot(libvirt.DOMAIN_REBOOT_ACPI_POWER_BTN)
	if err != nil {
		return fmt.Errorf("cannot reboot vm: %w", err)
	}
	return nil
}

func (c *QEMUConnection) Status(vm VM) (VMStatus, error) {
	var status VMStatus
	if c.conn == nil {
		return status, ErrConnectionIsNotReady
	}
	dom, err := c.conn.LookupDomainByName(vm.Name)
	if err != nil {
		return VMStatus{}, fmt.Errorf("cannot find vm by name: %w", err)
	}
	status.Active, err = dom.IsActive()
	if err != nil {
		return VMStatus{}, fmt.Errorf("cannot get vm status: %w", err)
	}
	return status, nil
}
