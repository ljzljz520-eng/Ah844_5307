package schema

import "logisticsconvert/model"

var SupportedDatabases = []string{"mysql", "postgres", "oracle", "sqlserver"}
var RequiredTables = []string{"vehicles", "drivers", "routes", "waybills"}

func IsSupportedDatabase(name string) bool {
	for _, v := range SupportedDatabases {
		if v == name {
			return true
		}
	}
	return false
}
func IsRequiredTable(name string) bool {
	for _, v := range RequiredTables {
		if v == name {
			return true
		}
	}
	return false
}
func NormalizeType(typ string) string {
	switch typ {
	case "INT", "INTEGER", "NUMBER":
		return "INTEGER"
	case "VARCHAR", "VARCHAR2", "TEXT":
		return "TEXT"
	case "DATE", "TIMESTAMP":
		return "TEXT"
	default:
		return "TEXT"
	}
}
func NormalizeTable(t model.Table) model.Table {
	out := t
	for i, v := range out.Types {
		out.Types[i] = NormalizeType(v)
	}
	return out
}
func SelectTables(src model.SourceSchema, names []string) model.SourceSchema {
	out := model.SourceSchema{Database: src.Database}
	for _, n := range names {
		if t, ok := src.Find(n); ok {
			out.Tables = append(out.Tables, NormalizeTable(t))
		}
	}
	return out
}
func MissingRequired(src model.SourceSchema) []string {
	var out []string
	for _, n := range RequiredTables {
		if _, ok := src.Find(n); !ok {
			out = append(out, n)
		}
	}
	return out
}
func TableSignature(t model.Table) string {
	s := t.Name + ":"
	for i, c := range t.Columns {
		s += c + "=" + NormalizeType(t.Types[i]) + ";"
	}
	return s
}
