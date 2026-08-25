package schema

import (
	"fmt"
	"logisticsconvert/model"
	"strings"
)

func ValidateSelection(src model.SourceSchema, names []string) error {
	if len(names) == 0 {
		return fmt.Errorf("select at least one table")
	}
	for _, n := range names {
		if _, ok := src.Find(n); !ok {
			return fmt.Errorf("table %s missing", n)
		}
	}
	return nil
}
func CanonicalDatabase(name string) string {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "mysql", "mariadb":
		return "mysql"
	case "postgres", "postgresql":
		return "postgres"
	case "oracle":
		return "oracle"
	case "sqlserver", "mssql":
		return "sqlserver"
	default:
		return ""
	}
}
func EnsurePrimaryKey(t model.Table) model.Table {
	if t.PrimaryKey != "" {
		return t
	}
	if len(t.Columns) > 0 {
		t.PrimaryKey = t.Columns[0]
	}
	return t
}
func ApplyDefaults(s model.SourceSchema) model.SourceSchema {
	for i := range s.Tables {
		s.Tables[i] = EnsurePrimaryKey(s.Tables[i])
		for j := range s.Tables[i].Types {
			if s.Tables[i].Types[j] == "" {
				s.Tables[i].Types[j] = "TEXT"
			}
		}
	}
	return s
}
func ValidateColumnNames(t model.Table) error {
	seen := map[string]bool{}
	for _, c := range t.Columns {
		if strings.TrimSpace(c) == "" {
			return fmt.Errorf("empty column")
		}
		if seen[c] {
			return fmt.Errorf("duplicate column %s", c)
		}
		seen[c] = true
	}
	return nil
}
func ValidateAll(s model.SourceSchema) error {
	if !IsSupportedDatabase(CanonicalDatabase(s.Database)) {
		return fmt.Errorf("unsupported database")
	}
	for _, t := range s.Tables {
		if err := ValidateColumnNames(t); err != nil {
			return err
		}
	}
	return nil
}
func Compatible(a, b model.Table) bool {
	if len(a.Columns) != len(b.Columns) {
		return false
	}
	for i := range a.Columns {
		if a.Columns[i] != b.Columns[i] || NormalizeType(a.Types[i]) != NormalizeType(b.Types[i]) {
			return false
		}
	}
	return true
}
