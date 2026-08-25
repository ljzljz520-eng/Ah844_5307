package store

import (
	"logisticsconvert/model"
	"sort"
)

func ByState(items []model.Conversion, state string) []model.Conversion {
	out := []model.Conversion{}
	for _, c := range items {
		if c.State == state {
			out = append(out, c)
		}
	}
	return out
}
func SortByID(items []model.Conversion) []model.Conversion {
	out := append([]model.Conversion(nil), items...)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func FindByDatabase(items []model.Conversion, name string) []model.Conversion {
	out := []model.Conversion{}
	for _, c := range items {
		if c.Source.Database == name {
			out = append(out, c)
		}
	}
	return out
}
func Latest(items []model.Conversion) model.Conversion {
	if len(items) == 0 {
		return model.Conversion{}
	}
	return SortByID(items)[len(items)-1]
}
func CountState(items []model.Conversion) map[string]int {
	out := map[string]int{}
	for _, c := range items {
		out[c.State]++
	}
	return out
}
func IDs(items []model.Conversion) []string {
	out := []string{}
	for _, c := range items {
		out = append(out, c.ID)
	}
	return out
}
