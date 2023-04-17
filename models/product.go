// Package models declare something
// MarsDong 2023/4/13
package models

// Product table structure
type Product struct {
	Uuid string `json:"Uuid"`
	Name string `json:"Name"`
	GormModel
}

// TableName dynamically returns the table name
// @Author MarsDong 2023-04-03 11:16:59
func (t *Product) TableName() string {
	return "product"
}

// GetUnique return record unique identify
// @Author MarsDong 2023-04-12 16:52:26
func (t *Product) GetUnique() (filed string, value interface{}) {
	return "uuid", t.Uuid
}

// Columns return all column names
// @Author MarsDong 2023-04-12 16:52:42
func (t *Product) Columns() []string {
	return t.GetColumns("uuid", "name")
}

// SummaryColumns return all summary column names.
// Reduce transmission loss when users retrieve large amounts of data
// @Author MarsDong 2023-04-12 16:52:58
func (t *Product) SummaryColumns() []string {
	return t.GetColumns("uuid", "name")
}
