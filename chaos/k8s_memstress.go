// Package chaos declare something
// MarsDong 2023/4/3
package chaos

import (
	"fmt"
)

const ChaosKindK8sMemStress string = "k8s.stress.mem"

type K8sMemStress struct {
	Meta
	Spec   *K8sMemStressSpec
	Status *K8sMemStressStatus
}

type K8sMemStressSpec struct {
	Percent  int    `json:"Percent"`
	Size     string `json:"Size"`
	Duration string `json:"Duration"`
	Spec
}

type K8sMemStressStatus struct {
	TaskID string `json:"TaskId"`
	CommonStatus
}

func (c *K8sMemStress) GetSelector() *Selector {
	return c.Spec.Selector
}

func (c *K8sMemStress) GetStatus() *CommonStatus {
	if c.Status == nil {
		c.Status = &K8sMemStressStatus{
			CommonStatus: CommonStatus{
				Status: "Init",
			},
		}
	}
	return &c.Status.CommonStatus
}

func (c *K8sMemStress) SetStatus(in interface{}) error {
	v, ok := in.(*K8sMemStressStatus)
	if !ok {
		return fmt.Errorf("invalid status format for K8sMemStress")
	}
	c.Status = v
	return nil
}

func init() {
	all.Register(ChaosKindK8sMemStress, new(K8sMemStress))
}
