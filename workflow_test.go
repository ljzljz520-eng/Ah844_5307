package main

import (
	"logisticsconvert/model"
	"logisticsconvert/service"
	"logisticsconvert/store"
	"os"
	"testing"
)

func fixture(t *testing.T) (*service.Service, func()) {
	f := t.TempDir() + "/x.db"
	s, e := store.Open(f)
	if e != nil {
		t.Fatal(e)
	}
	return service.New(s), func() { s.Close() }
}
func source() model.SourceSchema {
	return model.SourceSchema{Database: "mysql", Tables: []model.Table{{Name: "vehicles", Columns: []string{"id"}, Types: []string{"INT"}}, {Name: "drivers", Columns: []string{"id"}, Types: []string{"INT"}}, {Name: "routes", Columns: []string{"id"}, Types: []string{"INT"}}, {Name: "waybills", Columns: []string{"id"}, Types: []string{"INT"}}}}
}
func TestWorkflowOne(t *testing.T) {
	s, c := fixture(t)
	defer c()
	x, e := s.Create("one", source(), []string{"vehicles", "drivers"})
	if e != nil || x.State != "draft" {
		t.Fatal(e, x.State)
	}
	if _, e = s.Confirm("one"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, c := fixture(t)
	defer c()
	_, e := s.Create("two", source(), []string{"routes", "waybills"})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.Execute("two", true); e == nil {
		t.Fatal("confirmation bypass")
	}
}
func TestWorkflowThree(t *testing.T) {
	s, c := fixture(t)
	defer c()
	_, e := s.Create("three", source(), []string{"vehicles", "drivers", "routes", "waybills"})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.Confirm("three"); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Execute("three", true); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain22(t *testing.T) {
	s, c := fixture(t)
	defer c()
	_, e := s.Create("bug", source(), []string{"vehicles"})
	if e != nil {
		t.Fatal(e)
	}
	if e = s.Transition("bug", "executed"); e == nil {
		t.Fatal("invalid transition must be rejected")
	}
}
func TestEntryPointData(t *testing.T) {
	if _, e := os.Stat("go.mod"); e != nil {
		t.Fatal(e)
	}
}
