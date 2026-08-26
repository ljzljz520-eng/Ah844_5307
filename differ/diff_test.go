package differ

import (
	"logisticsconvert/model"
	"testing"
)

func TestBuildDDL(t *testing.T) {
	s := model.SourceSchema{Database: "mysql", Tables: []model.Table{{Name: "x", Columns: []string{"id"}, Types: []string{"INT"}}}}
	if len(BuildDDL(s).SQL) == 0 {
		t.Fatal()
	}
}
