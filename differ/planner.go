package differ

import (
	"fmt"
	"logisticsconvert/model"
	"logisticsconvert/schema"
)

type PlanStep struct {
	Order       int
	Action      string
	Table       string
	Destructive bool
}

func Plan(source, target model.SourceSchema) []PlanStep {
	ds := Compare(source, target)
	out := []PlanStep{}
	for i, d := range ds {
		des := d.Kind == "removed"
		out = append(out, PlanStep{Order: i + 1, Action: d.Kind, Table: d.Table, Destructive: des})
	}
	return out
}
func PlanText(p []PlanStep) string {
	out := ""
	for _, x := range p {
		flag := "safe"
		if x.Destructive {
			flag = "destructive"
		}
		out += fmt.Sprintf("%d. %s %s (%s)\n", x.Order, x.Action, x.Table, flag)
	}
	return out
}
func NormalizeTarget(t model.Table) model.Table {
	t = schema.NormalizeTable(t)
	return schema.EnsurePrimaryKey(t)
}
func Reconcile(source, target model.SourceSchema) model.SourceSchema {
	out := target
	for i := range out.Tables {
		out.Tables[i] = NormalizeTarget(out.Tables[i])
	}
	for _, t := range source.Tables {
		if _, ok := out.Find(t.Name); !ok {
			out.Tables = append(out.Tables, NormalizeTarget(t))
		}
	}
	return out
}
func SafePlan(source, target model.SourceSchema) []PlanStep {
	p := Plan(source, target)
	out := []PlanStep{}
	for _, x := range p {
		if !x.Destructive {
			out = append(out, x)
		}
	}
	return out
}
