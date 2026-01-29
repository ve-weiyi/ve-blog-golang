package dbparser

var commonMysqlDataTypeMapString = map[string]string{
	// For consistency, all integer types are converted to int64
	// bool
	"bool":    "bool",
	"boolean": "bool",
	// number
	"tinyint":   "int64",
	"smallint":  "int64",
	"mediumint": "int64",
	"int":       "int64",
	"int1":      "int64",
	"int2":      "int64",
	"int3":      "int64",
	"int4":      "int64",
	"int8":      "int64",
	"integer":   "int64",
	"bigint":    "int64",
	"float":     "float64",
	"float4":    "float64",
	"float8":    "float64",
	"double":    "float64",
	"decimal":   "float64",
	"dec":       "float64",
	"fixed":     "float64",
	"real":      "float64",
	"bit":       "byte",
	// date & time
	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"time":      "string",
	"year":      "int64",
	// string
	"linestring":      "string",
	"multilinestring": "string",
	"nvarchar":        "string",
	"nchar":           "string",
	"char":            "string",
	"bpchar":          "string",
	"character":       "string",
	"varchar":         "string",
	"binary":          "string",
	"bytea":           "string",
	"longvarbinary":   "string",
	"varbinary":       "string",
	"tinytext":        "string",
	"text":            "string",
	"mediumtext":      "string",
	"longtext":        "string",
	"enum":            "string",
	"set":             "string",
	"json":            "string",
	"jsonb":           "string",
	"blob":            "string",
	"longblob":        "string",
	"mediumblob":      "string",
	"tinyblob":        "string",
	"ltree":           "[]byte",
}

func ConvertMysqlToGoType(name string) string {
	return commonMysqlDataTypeMapString[name]
}

// convertNullTypeToPointer 将 sql.Null* 类型转换为对应的指针类型
func convertNullTypeToPointer(dataType string) string {
	switch dataType {
	case "sql.NullString":
		return "*string"
	case "sql.NullInt64":
		return "*int64"
	case "sql.NullInt32":
		return "*int32"
	case "sql.NullFloat64":
		return "*float64"
	case "sql.NullBool":
		return "*bool"
	case "sql.NullTime":
		return "*time.Time"
	default:
		return dataType
	}
}
