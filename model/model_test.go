package model

import "testing"

func TestEntitiesValidate(t *testing.T) {
	if (Vehicle{Status: "available", Capacity: 4}).IsAvailable() == false {
		t.Fatal()
	}
	if (Driver{Active: true, License: "x"}).IsEligible() == false {
		t.Fatal()
	}
	if (Route{Origin: "a", Destination: "b", Distance: 2}).IsUsable() == false {
		t.Fatal()
	}
}
