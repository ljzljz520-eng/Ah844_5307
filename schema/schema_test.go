package schema

import (
	"logisticsconvert/model"
	"testing"
)

func TestSchemaNormalize(t *testing.T) {
	if NormalizeType("VARCHAR") != "TEXT" {
		t.Fatal()
	}
	s := model.SourceSchema{Database: "mysql", Tables: []model.Table{{Name: "vehicles", Columns: []string{"id"}, Types: []string{"INT"}}}}
	if len(MissingRequired(s)) != 3 {
		t.Fatal()
	}
}
