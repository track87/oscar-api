// Package models declare something
// MarsDong 2023/4/13
package models

type TemplateProductView struct {
	Template
	ProductUuid string `json:"ProductUuid"`
	ProductName string `json:"ProductName"`
}

type TemplateProduct struct {
	GormModel
	TemplateUuid string `json:"TemplateUuid"`
	ProductUuid  string `json:"ProductUuid"`
}

// TableName dynamically returns the table name
// @Author MarsDong 2023-04-03 11:16:59
func (t *TemplateProduct) TableName() string {
	return "template_product"
}

// GetUnique return record unique identify
// @Author MarsDong 2023-04-12 16:52:26
func (t *TemplateProduct) GetUnique() (filed string, value interface{}) {
	return "id", t.ID
}

// Columns return all column names
// @Author MarsDong 2023-04-12 16:52:42
func (t *TemplateProduct) Columns() []string {
	return t.GetColumns("template_uuid", "product_uuid")
}

// SummaryColumns return all summary column names.
// Reduce transmission loss when users retrieve large amounts of data
// @Author MarsDong 2023-04-12 16:52:58
func (t *TemplateProduct) SummaryColumns() []string {
	return t.GetColumns("template_uuid", "product_uuid")
}
