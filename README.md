# Simple CLI for manipulation VMs through libvirt-daemon

## Tool Guide

This tool is supposed to be used as light-weight and easy utility for virtual machines (VM) manipulations.

### VM Creation

Create a machine "asgard" with defaults mem, disk and cpu. Machine name must be unique across your host OS.

```
$ libvirter create asgard
asgard created
```

Create a machine "noatun" with 2 CPU, 3 Gb drive and 100 Mb of memory use

```
$ libvirter create -c 2 -m 100 -d 3 noatun
noatun created
```

Create a "hell" machine using predefined sizes use option `-s` (size):

```
$ libvirter create -s tiny hell
hell created
```

Create a QEMU machine "jotunheim" use option `-t` (type):

```
$ libvirter create -t qemu jotunheim
jotunheim created
```

Create a KVM machine with a random name:

```
$ libvirter create -t kvm
asdfadf created
```

### VM Status

To get "asgard" machine status call:

```
$ libvirter status asgard
asgard is running
```

For a not running machine output will be:

```
asgard is stopped
```

If a VM does not exists you will receive following output and non-zero return code:

```
asgard not found
```

To get more information about a machine use option `-a`: 

```
$ libvirter status -a asgard
name: asgard
state: running
type: qemu
cpu: 1
memory: 512 Mb
disk: 3 Gb
uptime: 1h 12m 13s
```

### VM List (TBD)

To get VMs' list in a host OS call:

```
$ libvirter list
asgard
hell
noatun
```

Also, it's possible to provide a verbose list:

```
$ libvirter list -v
name	type	state	cpu	memory	disk	uptime
asgard	qemu	running	1	512	3	1h12m12s
noatun	kvm	stopped	2	1024	5	
hell	qemu	stopped	1	512	1	
```

*Note:* Some other options can be disscussed later if needed.

### VM Deletion

To delete "asgard" machine call:

```
$ liberter delete asgard
asgard deleted
```

### VM Starting

To start "asgard" machine call:

```
$ libvirter start asgard
asgard started
```

### VM Stopping

To stop "asgard" machine call:

```
$ libvirter stop asgard
asgard stopped
```

### VM Restarting

To restart "asgard" machine call:

```
$ libvirter restart asgard
asgard restarted
```

## Development

### Dependencies

#### Ubuntu

```bash
sudo apt-get install libvirt-dev
```

#### CentOS/Fedora
```bash
sudo dnf install libvirt-devel
```

### Build project

```bash
make build
```

Result will be in **bin** folder
