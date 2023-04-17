// Package chaosv1 declare something
// MarsDong 2023/3/31
package chaos

import (
	"fmt"
)

const ChaosKindHostMemStress string = "host.stress.mem"

type HostMemStress struct {
	Meta
	Spec   *HostMemStressSpec
	Status *HostMemStressStatus
}

type HostMemStressSpec struct {
	Percent  int    `json:"Percent"`
	Size     string `json:"Size"`
	Duration string `json:"Duration"`
	Spec
}

type HostMemStressStatus struct {
	TaskID string `json:"TaskId"`
	CommonStatus
}

func (c *HostMemStress) GetSelector() *Selector {
	return c.Spec.Selector
}

func (c *HostMemStress) GetStatus() *CommonStatus {
	if c.Status == nil {
		c.Status = &HostMemStressStatus{
			CommonStatus: CommonStatus{
				Status: "Init",
			},
		}
	}
	return &c.Status.CommonStatus
}

func (c *HostMemStress) SetStatus(in interface{}) error {
	v, ok := in.(*HostMemStressStatus)
	if !ok {
		return fmt.Errorf("invalid status format for HostMemStress")
	}
	c.Status = v
	return nil
}

func init() {
	all.Register(ChaosKindHostMemStress, new(HostMemStress))
}
