package main

import "testing"

func TestMainFixture(t *testing.T) {
	if "mysql" != "mysql" {
		t.Fatal()
	}
}
