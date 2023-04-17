// Package http declare something
// MarsDong 2023/3/31
package http

import (
	api "github.com/track87/oscar-api/workflow"
)

const (
	ActionCreateTemplate = "CreateTemplate"
	ActionDeleteTemplate = "DeleteTemplate"
)

type TemplateFilter struct {
	// delete by products
	Products []string `json:"Products"`
	// delete by names
	Names []string `json:"Names"`
	// delete by uuids
	Uuids []string `json:"Uuids"`
	// delete by creators
	Creator []string `json:"Creator"`
	Options string
}

// CreateTemplateReq defines request for a new template
type CreateTemplateReq struct {
	// Declare the association with the product
	Product string `json:"Product"`
	// Define other associated properties as required
	SomethingElse string        `json:"SomethingElse"`
	Template      *api.Workflow `json:"Template"`
	Common
}

// DeleteTemplateReq defines request for delete templates
// All conditions are intersected
type DeleteTemplateReq struct {
	TemplateFilter
	Common
}

// UpdateTemplateReq defines request for update special template
// The name is not updated
type UpdateTemplateReq struct {
	Template *api.Workflow `json:"Template"`
	Common
}

// ListTemplateReq defines filter condition
// All conditions are intersected
type ListTemplateReq struct {
	TemplateFilter
	ListReq
}

func init() {
	all.Register(ActionCreateTemplate, new(CreateTemplateReq))
	all.Register(ActionDeleteTemplate, new(DeleteTemplateReq))
}
