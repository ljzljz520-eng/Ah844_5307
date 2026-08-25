package parser

import (
	"fmt"
	"logisticsconvert/model"
	"strings"
)

func ParseCreateTable(sql string) (model.Table, error) {
	clean := strings.TrimSpace(sql)
	upper := strings.ToUpper(clean)
	if !strings.HasPrefix(upper, "CREATE TABLE") {
		return model.Table{}, fmt.Errorf("not create table")
	}
	open := strings.Index(clean, "(")
	close := strings.LastIndex(clean, ")")
	if open < 0 || close < open {
		return model.Table{}, fmt.Errorf("malformed ddl")
	}
	head := strings.TrimSpace(clean[len("CREATE TABLE"):open])
	if head == "" {
		return model.Table{}, fmt.Errorf("missing name")
	}
	parts := strings.Split(clean[open+1:close], ",")
	t := model.Table{Name: head}
	for _, p := range parts {
		f := strings.Fields(strings.TrimSpace(p))
		if len(f) >= 2 {
			t.Columns = append(t.Columns, f[0])
			t.Types = append(t.Types, f[1])
		}
	}
	return t, t.Validate()
}
func SplitStatements(sql string) []string {
	raw := strings.Split(sql, ";")
	out := []string{}
	for _, s := range raw {
		if strings.TrimSpace(s) != "" {
			out = append(out, strings.TrimSpace(s))
		}
	}
	return out
}
func QuoteIdentifier(name string) string { return "\"" + strings.ReplaceAll(name, "\"", "\"\"") + "\"" }
func BuildInsert(t model.Table, values []string) string {
	cols := strings.Join(t.Columns, ", ")
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", QuoteIdentifier(t.Name), cols, strings.Join(values, ", "))
}
func RenderType(typ string, database string) string {
	if database == "oracle" && typ == "TEXT" {
		return "VARCHAR2(255)"
	}
	if database == "mysql" && typ == "TEXT" {
		return "VARCHAR(255)"
	}
	return typ
}
func NormalizeSQL(sql string) string { return strings.TrimSpace(strings.ReplaceAll(sql, "\r\n", "\n")) }
