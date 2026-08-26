package differ

import (
	"fmt"
	"logisticsconvert/model"
	"logisticsconvert/schema"
	"sort"
)

func Compare(source, target model.SourceSchema) []model.Difference {
	var out []model.Difference
	for _, st := range source.Tables {
		tt, ok := target.Find(st.Name)
		if !ok {
			out = append(out, model.Difference{Table: st.Name, Kind: "removed", Detail: "source table absent in target"})
			continue
		}
		out = append(out, compareTable(st, tt)...)
	}
	for _, tt := range target.Tables {
		if _, ok := source.Find(tt.Name); !ok {
			out = append(out, model.Difference{Table: tt.Name, Kind: "added", Detail: "target-only table"})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Table < out[j].Table })
	return out
}
func compareTable(a, b model.Table) []model.Difference {
	var out []model.Difference
	if schema.TableSignature(a) != schema.TableSignature(b) {
		out = append(out, model.Difference{Table: a.Name, Kind: "changed", Detail: fmt.Sprintf("%s -> %s", schema.TableSignature(a), schema.TableSignature(b))})
	}
	return out
}
func BuildDDL(s model.SourceSchema) model.TargetDDL {
	var sql string
	for _, t := range s.Tables {
		sql += "CREATE TABLE " + t.Name + " ("
		for i, c := range t.Columns {
			if i > 0 {
				sql += ", "
			}
			sql += c + " " + schema.NormalizeType(t.Types[i])
		}
		sql += ");\n"
	}
	return model.TargetDDL{SQL: sql, Tables: s.Tables}
}
func Report(ds []model.Difference) string {
	if len(ds) == 0 {
		return "no differences"
	}
	out := ""
	for _, d := range ds {
		out += d.Kind + " " + d.Table + ": " + d.Detail + "\n"
	}
	return out
}
func HasBlocking(ds []model.Difference) bool {
	for _, d := range ds {
		if d.Kind == "removed" {
			return true
		}
	}
	return false
}
