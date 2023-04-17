// Package models declare something
// MarsDong 2023/4/17
package models

import (
	"time"
)

// GormModel gorm base model
type GormModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Creator   string
	Updater   string
}

func (m *GormModel) GetColumns(extraColumns ...string) []string {
	columns := []string{"id", "created_at", "updated_at", "creator", "updater"}
	columns = append(columns, extraColumns...)
	return columns
}

// GetID return table id
// @Author MarsDong 2022-11-18 15:02:15
func (m *GormModel) GetID() uint {
	return m.ID
}

// ResetID set table is with zero
// @Author MarsDong 2022-11-18 15:02:22
func (m *GormModel) ResetID() {
	m.ID = 0
}
