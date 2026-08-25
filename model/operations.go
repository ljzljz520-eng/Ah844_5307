package model

import "sort"

func SortTables(t []Table) []Table {
	out := append([]Table(nil), t...)
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
func TableNames(s SourceSchema) []string {
	out := make([]string, 0, len(s.Tables))
	for _, t := range s.Tables {
		out = append(out, t.Name)
	}
	sort.Strings(out)
	return out
}
func FilterTables(s SourceSchema, prefix string) SourceSchema {
	out := SourceSchema{Database: s.Database}
	for _, t := range s.Tables {
		if len(prefix) == 0 || len(t.Name) >= len(prefix) && t.Name[:len(prefix)] == prefix {
			out.Tables = append(out.Tables, t)
		}
	}
	return out
}
func AddColumn(t Table, name, typ string) Table {
	t.Columns = append(t.Columns, name)
	t.Types = append(t.Types, typ)
	return t
}
func RemoveColumn(t Table, name string) Table {
	out := Table{Name: t.Name, PrimaryKey: t.PrimaryKey}
	for i, c := range t.Columns {
		if c != name {
			out.Columns = append(out.Columns, c)
			out.Types = append(out.Types, t.Types[i])
		}
	}
	return out
}
func RenameTable(s SourceSchema, from, to string) SourceSchema {
	for i := range s.Tables {
		if s.Tables[i].Name == from {
			s.Tables[i].Name = to
		}
	}
	return s
}
func MergeSchemas(a, b SourceSchema) SourceSchema {
	out := SourceSchema{Database: a.Database, Tables: append([]Table(nil), a.Tables...)}
	for _, t := range b.Tables {
		found := false
		for _, x := range out.Tables {
			if x.Name == t.Name {
				found = true
			}
		}
		if !found {
			out.Tables = append(out.Tables, t)
		}
	}
	return out
}
func ConversionSummary(c Conversion) string {
	return c.ID + "/" + c.State + "/" + string(rune(len(c.Source.Tables)+'0'))
}
func (c *Conversion) MarkReviewed() {
	if c.State == "draft" {
		c.State = "reviewed"
	}
}
func (c Conversion) IsTerminal() bool               { return c.State == "executed" || c.State == "failed" }
func (c Conversion) DifferenceCount() int           { return len(c.Differences) }
func (c Conversion) TableSelected(name string) bool { _, ok := c.Source.Find(name); return ok }
