package service

import (
	"logisticsconvert/model"
	"logisticsconvert/store"
	"testing"
)

func TestConfirmExecute(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	svc := New(s)
	src := model.SourceSchema{Database: "mysql", Tables: []model.Table{{Name: "x", Columns: []string{"id"}, Types: []string{"INT"}}}}
	if _, e = svc.Create("x", src, []string{"x"}); e != nil {
		t.Fatal(e)
	}
	if _, e = svc.Confirm("x"); e != nil {
		t.Fatal(e)
	}
	if _, e = svc.Execute("x", true); e != nil {
		t.Fatal(e)
	}
}
