package store

import (
	"logisticsconvert/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.PutConversion(model.Conversion{ID: "persist", State: "draft"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetConversion("persist"); e != nil {
		t.Fatal(e)
	}
}
