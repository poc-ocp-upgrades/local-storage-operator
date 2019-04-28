package diskmaker

import (
	"testing"
)

func TestFindMatchingDisk(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := NewDiskMaker("/tmp/foo", "/mnt/local-storage")
	deviceSet, err := d.findNewDisks(getData())
	if err != nil {
		t.Fatalf("error getting data %v", err)
	}
	if len(deviceSet) != 7 {
		t.Errorf("expected 7 devices got %d", len(deviceSet))
	}
	diskConfig := map[string]*Disks{"foo": &Disks{DeviceIDs: []string{"xyz"}}}
	allDiskIds := getDeiveIDs()
	deviceMap, err := d.findMatchingDisks(diskConfig, deviceSet, allDiskIds)
	if err != nil {
		t.Fatalf("error finding matchin device %v", err)
	}
	if len(deviceMap) != 0 {
		t.Errorf("expected 0 elements in map got %d", len(deviceMap))
	}
}
func getData() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return `
sda
sda1 /boot
sda2 [SWAP]
sda3 /
vda
vdb
vdc
vdd
vde
vdf`
}
func getDeiveIDs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"/dev/disk/by-id/xyz"}
}
