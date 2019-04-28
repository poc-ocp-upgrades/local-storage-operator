package diskmaker

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"
)

var (
	checkDuration	= 5 * time.Second
	diskByIDPath	= "/dev/disk/by-id/*"
)

type DiskMaker struct {
	configLocation	string
	symlinkLocation	string
}
type DiskLocation struct {
	diskName	string
	diskID		string
}

func NewDiskMaker(configLocation, symLinkLocation string) *DiskMaker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := &DiskMaker{}
	t.configLocation = configLocation
	t.symlinkLocation = symLinkLocation
	return t
}
func (d *DiskMaker) loadConfig() (DiskConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	content, err := ioutil.ReadFile(d.configLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s with %v", d.configLocation, err)
	}
	var diskConfig DiskConfig
	err = yaml.Unmarshal(content, &diskConfig)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling %s with %v", d.configLocation, err)
	}
	return diskConfig, nil
}
func (d *DiskMaker) Run(stop <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ticker := time.NewTicker(checkDuration)
	defer ticker.Stop()
	err := os.MkdirAll(d.symlinkLocation, 0755)
	if err != nil {
		logrus.Errorf("error creating local-storage directory %s with %v", d.symlinkLocation, err)
		os.Exit(-1)
	}
	for {
		select {
		case <-ticker.C:
			diskConfig, err := d.loadConfig()
			if err != nil {
				logrus.Errorf("error loading configuration with %v", err)
				break
			}
			d.symLinkDisks(diskConfig)
		case <-stop:
			logrus.Infof("exiting, received message on stop channel")
			os.Exit(0)
		}
	}
}
func (d *DiskMaker) symLinkDisks(diskConfig DiskConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := exec.Command("lsblk", "--list", "-o", "NAME,MOUNTPOINT", "--noheadings")
	var out bytes.Buffer
	var err error
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("error running lsblk %v", err)
		return
	}
	deviceSet, err := d.findNewDisks(out.String())
	if err != nil {
		logrus.Errorf("error unmrashalling json %v", err)
		return
	}
	if len(deviceSet) == 0 {
		logrus.Infof("unable to find any new disks")
		return
	}
	allDiskIds, err := filepath.Glob(diskByIDPath)
	if err != nil {
		logrus.Errorf("error listing disks in /dev/disk/by-id : %v", err)
		return
	}
	deviceMap, err := d.findMatchingDisks(diskConfig, deviceSet, allDiskIds)
	if err != nil {
		logrus.Errorf("error matching finding disks : %v", err)
		return
	}
	if len(deviceMap) == 0 {
		logrus.Errorf("unable to find any matching disks")
		return
	}
	for storageClass, deviceArray := range deviceMap {
		for _, deviceNameLoction := range deviceArray {
			symLinkDirPath := path.Join(d.symlinkLocation, storageClass)
			err := os.MkdirAll(symLinkDirPath, 0755)
			if err != nil {
				logrus.Errorf("error creating symlink directory %s with %v", symLinkDirPath, err)
				continue
			}
			symLinkPath := path.Join(symLinkDirPath, deviceNameLoction.diskName)
			var symLinkErr error
			if deviceNameLoction.diskID != "" {
				logrus.Infof("symlinking to %s to %s", deviceNameLoction.diskID, symLinkPath)
				symLinkErr = os.Symlink(deviceNameLoction.diskID, symLinkPath)
			} else {
				devicePath := path.Join("/dev", deviceNameLoction.diskName)
				logrus.Infof("symlinking to %s to %s", devicePath, symLinkPath)
				symLinkErr = os.Symlink(devicePath, symLinkPath)
			}
			if symLinkErr != nil {
				logrus.Errorf("error creating symlink %s with %v", symLinkPath, err)
			}
		}
	}
}
func (d *DiskMaker) findMatchingDisks(diskConfig DiskConfig, deviceSet sets.String, allDiskIds []string) (map[string][]DiskLocation, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	blockDeviceMap := make(map[string][]DiskLocation)
	addDiskToMap := func(scName, stableDeviceID, diskName string) {
		deviceArray, ok := blockDeviceMap[scName]
		if !ok {
			deviceArray = []DiskLocation{}
		}
		deviceArray = append(deviceArray, DiskLocation{diskName, stableDeviceID})
		blockDeviceMap[scName] = deviceArray
	}
	for storageClass, disks := range diskConfig {
		for _, diskName := range disks.DiskNames {
			if hasExactDisk(deviceSet, diskName) {
				matchedDeviceID, err := d.findStableDeviceID(diskName, allDiskIds)
				if err != nil {
					logrus.Errorf("Unable to find disk ID %s for local pool %v", diskName, err)
					addDiskToMap(storageClass, "", diskName)
					continue
				}
				addDiskToMap(storageClass, matchedDeviceID, diskName)
				continue
			}
		}
		for _, deviceID := range disks.DeviceIDs {
			matchedDeviceID, matchedDiskName, err := d.findDeviceByID(deviceID)
			if err != nil {
				logrus.Errorf("unable to add disk-id %s to local disk pool %v", deviceID, err)
				continue
			}
			addDiskToMap(storageClass, matchedDeviceID, matchedDiskName)
		}
	}
	return blockDeviceMap, nil
}
func (d *DiskMaker) findDeviceByID(deviceID string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	completeDiskIDPath := fmt.Sprintf("%s/%s", diskByIDPath, deviceID)
	diskDevPath, err := filepath.EvalSymlinks(completeDiskIDPath)
	if err != nil {
		return "", "", fmt.Errorf("unable to find device with id %s", deviceID)
	}
	diskDevName := filepath.Base(diskDevPath)
	return completeDiskIDPath, diskDevName, nil
}
func (d *DiskMaker) findStableDeviceID(diskName string, allDisks []string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, diskIDPath := range allDisks {
		diskDevPath, err := filepath.EvalSymlinks(diskIDPath)
		if err != nil {
			continue
		}
		diskDevName := filepath.Base(diskDevPath)
		if diskDevName == diskName {
			return diskIDPath, nil
		}
	}
	return "", fmt.Errorf("unable to find ID of disk %s", diskName)
}
func (d *DiskMaker) findNewDisks(content string) (sets.String, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	deviceSet := sets.NewString()
	deviceLines := strings.Split(content, "\n")
	for _, deviceLine := range deviceLines {
		deviceLine := strings.TrimSpace(deviceLine)
		deviceDetails := strings.Split(deviceLine, " ")
		if len(deviceDetails) == 1 && len(deviceDetails[0]) > 0 {
			deviceSet.Insert(deviceDetails[0])
		}
	}
	return deviceSet, nil
}
func hasExactDisk(disks sets.String, device string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, disk := range disks.List() {
		if disk == device {
			return true
		}
	}
	return false
}
