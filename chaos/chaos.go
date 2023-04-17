// Package chaos declare something
// MarsDong 2023/3/30
package chaos

import (
	"time"
)

const (
	StageDrill   = "drill"
	StageRecover = "Recover"
)

type IChaos interface {
	GetMeta() Meta
	GetSelector() *Selector
	GetStatus() *CommonStatus
	SetStatus(in interface{}) error
}

type Meta struct {
	Kind string
	Name string
}

func (m Meta) GetMeta() Meta {
	return m
}

type Spec struct {
	Selector *Selector
}

type CommonStatus struct {
	Stage          string     `json:"Stage"`
	Status         string     `json:"Status"`
	Message        string     `json:"Message"`
	StartTime      *time.Time `json:"StartTime"`
	EndTime        *time.Time `json:"EndTime"`
	DrillRecords   []*Record  `json:"DrillRecords"`
	RecoverRecords []*Record  `json:"RecoverRecords"`
}

type Record struct {
	Zone       string     `json:"Zone"`
	Target     string     `json:"Target"`
	TargetType string     `json:"TargetType"`
	Status     string     `json:"Status"`
	Message    string     `json:"Message"`
	Detail     string     `json:"Detail"`
	StartTime  *time.Time `json:"StartTime"`
	EndTime    *time.Time `json:"EndTime"`
}
