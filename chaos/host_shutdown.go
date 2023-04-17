// Package chaos declare something
// MarsDong 2023/4/3
package chaos

import (
	"fmt"
)

const ChaosKindHostShutdown string = "host.shutdown"

type HostShutdown struct {
	Meta
	Spec   *HostShutdownSpec
	Status *HostShutdownStatus
}

type HostShutdownSpec struct {
	Spec
}

type HostShutdownStatus struct {
	TaskID string `json:"TaskId"`
	CommonStatus
}

func (c *HostShutdown) GetSelector() *Selector {
	return c.Spec.Selector
}

func (c *HostShutdown) GetStatus() *CommonStatus {
	if c.Status == nil {
		c.Status = &HostShutdownStatus{
			CommonStatus: CommonStatus{
				Status: "Init",
			},
		}
	}
	return &c.Status.CommonStatus
}

func (c *HostShutdown) SetStatus(in interface{}) error {
	v, ok := in.(*HostShutdownStatus)
	if !ok {
		return fmt.Errorf("invalid status format for HostShutdown")
	}
	c.Status = v
	return nil
}

func init() {
	all.Register(ChaosKindHostShutdown, new(HostShutdown))
}
