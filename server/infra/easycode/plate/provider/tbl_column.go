package provider

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/easycode/plate/field"
)

// 需要的数据
type Column struct {
	gorm.ColumnType
	TableName string `gorm:"column:TABLE_NAME"`
	Indexes   []*Index

	ColumnName      string `json:"columnName" gorm:"column:column_name"`       //列名
	ColumnFiledType string `json:"ColumnFiledType" gorm:"column:column_type"`  //字段类型 varchar(11)
	ColumnDefault   string `json:"columnDefault" gorm:"column:column_default"` //默认值
	ColumnComment   string `json:"columnComment" gorm:"column:column_comment"` //备注
	DataType        string `json:"dataType" gorm:"column:data_type"`           //数据类型 varchar
	DataTypeLong    int64  `json:"dataTypeLong" gorm:"column:data_type_long"`  //数据长度
	IsNullable      bool   `json:"isNullable" gorm:"column:is_nullable"`       //是否可空
	IsPrimaryKey    bool   //数据是否是主键
	IsUnique        bool   //数据是否是唯一的
	IsAutoIncrement bool   //数据是否是自动递增
	HasDefault      bool   //数据是否有默认值
}

// GetDataType get data type
func (c *Column) GetDataType(cfg *field.FieldConfig) (fieldtype string) {
	if mapping, ok := cfg.DataTypeMap[c.DatabaseTypeName()]; ok {
		return mapping(c.ColumnType)
	}
	if cfg.UseScanType && c.ScanType() != nil {
		return c.ScanType().String()
	}
	return dataType.Get(c.DatabaseTypeName(), c.columnType())
}

// ToField convert to field
func (c *Column) ToField(cfg *field.FieldConfig) *Field {
	fieldType := c.GetDataType(cfg)
	if cfg.FieldSignable && strings.Contains(c.columnType(), "unsigned") && strings.HasPrefix(fieldType, "int") {
		fieldType = "u" + fieldType
	}
	switch {
	case c.Name() == "deleted_at" && fieldType == "time.Time":
		fieldType = "gorm.DeletedAt"
	case cfg.FieldCoverable && c.needDefaultTag(c.defaultTagValue()):
		fieldType = "*" + fieldType
	case cfg.FieldNullable:
		if n, ok := c.Nullable(); ok && n {
			fieldType = "*" + fieldType
		}
	}

	var comment string
	if c, ok := c.Comment(); ok {
		comment = c
	}

	//if token.IsKeyword(fileName) {
	//	fileName = fileName + "_"
	//}

	return &Field{
		Name:             c.Name(),
		FieldName:        cfg.FieldNameNS(c.ColumnName),
		FieldType:        fieldType,
		ColumnComment:    comment,
		MultilineComment: false,
		Tag:              map[string]string{field.TagKeyJson: cfg.FieldJSONTagNS(c.Name())},
		GORMTag:          c.buildGormTag(),
	}
}

func (c *Column) multilineComment() bool {
	cm, ok := c.Comment()
	return ok && strings.Contains(cm, "\n")
}

func (c *Column) buildGormTag() field.GormTag {
	tag := field.GormTag{
		field.TagKeyGormColumn: []string{c.Name()},
		field.TagKeyGormType:   []string{c.columnType()},
	}
	isPriKey, ok := c.PrimaryKey()
	isValidPriKey := ok && isPriKey
	if isValidPriKey {
		tag.Set(field.TagKeyGormPrimaryKey, "")
		if at, ok := c.AutoIncrement(); ok {
			tag.Set(field.TagKeyGormAutoIncrement, fmt.Sprintf("%t", at))
		}
	} else if n, ok := c.Nullable(); ok && !n {
		tag.Set(field.TagKeyGormNotNull, "")
	}

	for _, idx := range c.Indexes {
		if idx == nil {
			continue
		}
		if pk, _ := idx.PrimaryKey(); pk { //ignore PrimaryKey
			continue
		}
		if uniq, _ := idx.Unique(); uniq {
			tag.Append(field.TagKeyGormUniqueIndex, fmt.Sprintf("%s,priority:%d", idx.Name(), idx.Priority))
		} else {
			tag.Append(field.TagKeyGormIndex, fmt.Sprintf("%s,priority:%d", idx.Name(), idx.Priority))
		}
	}

	if dtValue := c.defaultTagValue(); c.needDefaultTag(dtValue) { // cannot set default tag for primary key
		tag.Set(field.TagKeyGormDefault, dtValue)
	}
	if comment, ok := c.Comment(); ok && comment != "" {
		if c.multilineComment() {
			comment = strings.ReplaceAll(comment, "\n", "\\n")
		}
		tag.Set(field.TagKeyGormComment, comment)
	}
	return tag
}

// needDefaultTag check if default tag needed
func (c *Column) needDefaultTag(defaultTagValue string) bool {
	if defaultTagValue == "" {
		return false
	}
	switch c.ScanType().Kind() {
	case reflect.Bool:
		return defaultTagValue != "false"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return defaultTagValue != "0"
	case reflect.String:
		return defaultTagValue != ""
	case reflect.Struct:
		return strings.Trim(defaultTagValue, "'0:- ") != ""
	}
	return c.Name() != "created_at" && c.Name() != "updated_at"
}

// defaultTagValue return gorm default tag's value
func (c *Column) defaultTagValue() string {
	value, ok := c.DefaultValue()
	if !ok {
		return ""
	}
	if value != "" && strings.TrimSpace(value) == "" {
		return "'" + value + "'"
	}
	return value
}

func (c *Column) columnType() (v string) {
	if cl, ok := c.ColumnType.ColumnType(); ok {
		return cl
	}
	return c.DatabaseTypeName()
}
