package differ

import (
	"logisticsconvert/model"
	"logisticsconvert/schema"
	"strings"
)

func RenderDDL(s model.SourceSchema, dialect string) string {
	var b strings.Builder
	for _, t := range s.Tables {
		b.WriteString("CREATE TABLE " + schema.QuoteIfNeeded(t.Name) + " (\n")
		for i, c := range t.Columns {
			if i > 0 {
				b.WriteString(",\n")
			}
			b.WriteString("  " + schema.QuoteIfNeeded(c) + " " + renderType(t.Types[i], dialect))
		}
		b.WriteString("\n);\n")
	}
	return b.String()
}
func renderType(t, d string) string {
	if d == "oracle" && schema.NormalizeType(t) == "TEXT" {
		return "VARCHAR2(255)"
	}
	if d == "mysql" && schema.NormalizeType(t) == "TEXT" {
		return "VARCHAR(255)"
	}
	return schema.NormalizeType(t)
}
func RenderReport(ds []model.Difference, format string) string {
	if format == "json" {
		return Report(ds)
	}
	return Report(ds)
}
func GroupByKind(ds []model.Difference) map[string][]model.Difference {
	out := map[string][]model.Difference{}
	for _, d := range ds {
		out[d.Kind] = append(out[d.Kind], d)
	}
	return out
}
func ChangedTables(ds []model.Difference) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, d := range ds {
		if d.Kind == "changed" && !seen[d.Table] {
			seen[d.Table] = true
			out = append(out, d.Table)
		}
	}
	return out
}
