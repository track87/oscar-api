// Package models declare something
// MarsDong 2023/4/3
package models

// Template table structure
type Template struct {
	Uuid string `json:"Uuid"`
	Name string `json:"Name"`
	GormModel
}

// TableName dynamically returns the table name
// @Author MarsDong 2023-04-03 11:16:59
func (t *Template) TableName() string {
	return "template"
}

// GetUnique return record unique identify
// @Author MarsDong 2023-04-12 16:52:26
func (t *Template) GetUnique() (filed string, value interface{}) {
	return "uuid", t.Uuid
}

// Columns return all column names
// @Author MarsDong 2023-04-12 16:52:42
func (t *Template) Columns() []string {
	return t.GetColumns("uuid", "name")
}

// SummaryColumns return all summary column names.
// Reduce transmission loss when users retrieve large amounts of data
// @Author MarsDong 2023-04-12 16:52:58
func (t *Template) SummaryColumns() []string {
	return t.GetColumns("uuid", "name")
}
