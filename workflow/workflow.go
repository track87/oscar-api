// Copyright 2023 MarsDong.
//
// Package v1 declare something
// MarsDong 2023/3/31

package workflow

import (
	"encoding/json"
	"time"

	"github.com/track87/oscar-api/chaos"
)

type NodeKind string

const (
	NodeKindWorkflow     = "Workflow"
	NodeKindWorkflowNode = "Node"
)

type WorkflowPhase string

const (
	WorkflowUnknown   WorkflowPhase = ""
	WorkflowPending   WorkflowPhase = "Pending" // pending some set-up - rarely used
	WorkflowRunning   WorkflowPhase = "Running" // any node has started; pods might not be running yet, the workflow maybe suspended too
	WorkflowSucceeded WorkflowPhase = "Succeeded"
	WorkflowFailed    WorkflowPhase = "Failed" // it maybe that the workflow was terminated
	WorkflowError     WorkflowPhase = "Error"
)

func (p WorkflowPhase) Completed() bool {
	switch p {
	case WorkflowSucceeded, WorkflowFailed, WorkflowError:
		return true
	default:
		return false
	}
}

// IWfNode defines common operation for workfow node
type IWfNode interface {
	GetEntry() string
	GetMeta() NodeMeta
	GetIncoming() []string
	GetOutgoing() []string
	GetStatus() *NodeStatus
	SetStatus(in *NodeStatus)
}

type NodeMeta struct {
	Uuid string   `json:"Uuid"`
	Name string   `json:"Name"`
	Kind NodeKind `json:"Kind"`
}

type NodeSpec struct {
	Entry       string
	WaitBefore  string `json:"WaitBefore"`
	WaitAfter   string `json:"WaitAfter"`
	AfterFailed string `json:"AfterFailed"`
	// FailFast, if specified, will fail this template if any of its child has failed
	FailFast *bool
	Incoming []string `json:"Incoming"`
	Outgoing []string `json:"Outgoing"`
}

type NodeStatus struct {
	Phase             WorkflowPhase `json:"Phase"`
	Message           string        `json:"Message"`
	StartAt           time.Time     `json:"StartAt"`
	FinishedAt        time.Time     `json:"FinishedAt"`
	EstimatedDuration int           `json:"EstimatedDuration"`
}

type Workflow struct {
	NodeMeta
	Spec   *WorkflowSpec
	Status *NodeStatus
}

type WorkflowSpec struct {
	NodesMap map[string]IWfNode
	NodeSpec
}

type WorkflowStatus struct {
	NodeStatus
}

func (w *Workflow) GetMeta() NodeMeta {
	return w.NodeMeta
}

func (w *Workflow) GetIncoming() []string {
	return w.Spec.Incoming
}

func (w *Workflow) GetOutgoing() []string {
	return w.Spec.Outgoing
}

func (w *Workflow) GetNodes() []IWfNode {
	nodes := make([]IWfNode, 0)
	for _, node := range w.Spec.NodesMap {
		nodes = append(nodes, node)
	}
	return nodes
}

func (w *Workflow) GetEntry() string {
	return w.Spec.Entry
}

func (w *Workflow) GetStatus() *NodeStatus {
	return w.Status
}

func (w *Workflow) SetStatus(in *NodeStatus) {
	w.Status = in
	return
}

type WfNode struct {
	NodeMeta
	Spec   *WfNodeSpec
	Status *NodeStatus
}

type WfNodeSpec struct {
	Chaos []chaos.IChaos `json:"Chaos"`
	NodeSpec
}

type WfNodeStatus struct {
	NodeStatus
}

func (n *WfNode) GetIncoming() []string {
	return n.Spec.Incoming
}

func (n *WfNode) GetOutgoing() []string {
	return n.Spec.Outgoing
}

func (n *WfNode) GetMeta() NodeMeta {
	return n.NodeMeta
}

func (n *WfNode) GetEntry() string {
	return n.Spec.Entry
}

func (n *WfNode) GetStatus() *NodeStatus {
	return n.Status
}

func (n *WfNode) SetStatus(in *NodeStatus) {
	n.Status = in
	return
}

func SpawnWfObject(kind NodeKind, in IWfNode) interface{} {
	objectBytes, _ := json.Marshal(in)
	switch kind {
	case NodeKindWorkflow:
		object := &Workflow{}
		_ = json.Unmarshal(objectBytes, object)
		return object
	case NodeKindWorkflowNode:
		object := &WfNode{}
		_ = json.Unmarshal(objectBytes, object)
		return object
	default:
		return nil
	}
}
