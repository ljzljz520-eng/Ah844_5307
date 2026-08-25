package parser

import (
	"encoding/json"
	"fmt"
	"logisticsconvert/model"
	"logisticsconvert/schema"
	"strings"
)

func ParseJSON(data []byte) (model.SourceSchema, error) {
	var s model.SourceSchema
	if err := json.Unmarshal(data, &s); err != nil {
		return s, err
	}
	return ParseSchema(s)
}
func ParseSchema(s model.SourceSchema) (model.SourceSchema, error) {
	if !schema.IsSupportedDatabase(strings.ToLower(s.Database)) {
		return s, fmt.Errorf("unsupported database %s", s.Database)
	}
	for i := range s.Tables {
		s.Tables[i] = schema.NormalizeTable(s.Tables[i])
	}
	if err := s.Validate(); err != nil {
		return s, err
	}
	return s, nil
}
func ParseColumns(name string, columns []string, types []string) (model.Table, error) {
	t := model.Table{Name: name, Columns: columns, Types: types}
	if err := t.Validate(); err != nil {
		return t, err
	}
	return schema.NormalizeTable(t), nil
}
func DatabaseHint(data []byte) string {
	var x struct {
		Database string `json:"database"`
	}
	_ = json.Unmarshal(data, &x)
	return strings.ToLower(x.Database)
}
func EncodeSchema(s model.SourceSchema) ([]byte, error) { return json.MarshalIndent(s, "", "  ") }
func DecodeTables(data []byte) ([]model.Table, error) {
	var s model.SourceSchema
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return s.Tables, nil
}
