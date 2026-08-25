package main

import (
	"logisticsconvert/model"
	"logisticsconvert/store"
	"testing"
)

func TestStoreLifecycle(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.PutConversion(model.Conversion{ID: "a"}); e != nil {
		t.Fatal(e)
	}
	if n, e := s.Count(); e != nil || n != 1 {
		t.Fatal(n, e)
	}
}
