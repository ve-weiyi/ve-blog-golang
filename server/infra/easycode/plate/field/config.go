package field

import "gorm.io/gorm"

type FieldConfig struct {
	//gorm.ColumnType
	//TableName   string                                                        `gorm:"column:TABLE_NAME"`
	//Indexes     []*Index                                                      `gorm:"-"`
	UseScanType                   bool                                                          `gorm:"-"`
	dataTypeMap                   map[string]func(columnType gorm.ColumnType) (dataType string) `gorm:"-"`
	jsonTagNS                     func(columnName string) string                                `gorm:"-"`
	nullable, coverable, signable bool

	DataTypeMap map[string]func(columnType gorm.ColumnType) (dataType string)

	FieldNullable     bool // generate pointer when field is nullable
	FieldCoverable    bool // generate pointer when field has default value
	FieldSignable     bool // detect integer field's unsigned type, adjust generated data type
	FieldWithIndexTag bool // generate with gorm index tag
	FieldWithTypeTag  bool // generate with gorm column type tag

	FieldNameNS    func(columnName string) string
	FieldJSONTagNS func(columnName string) string
}

// SetDataTypeMap set data type map
func (c *FieldConfig) SetDataTypeMap(m map[string]func(columnType gorm.ColumnType) (dataType string)) {
	c.dataTypeMap = m
}

// WithNS with name strategy
func (c *FieldConfig) WithNS(jsonTagNS func(columnName string) string) {
	c.jsonTagNS = jsonTagNS
	if c.jsonTagNS == nil {
		c.jsonTagNS = func(n string) string { return n }
	}
}
