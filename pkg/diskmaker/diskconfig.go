package diskmaker

import (
	"fmt"
	"github.com/ghodss/yaml"
)

type Disks struct {
	DiskNames	[]string	`json:"disks,omitempty"`
	DeviceIDs	[]string	`json:"deviceIDs,omitempty"`
}
type DiskConfig map[string]*Disks

func (d *DiskConfig) ToYAML() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	y, err := yaml.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("error marshaling to yaml: %v", err)
	}
	return string(y), nil
}
