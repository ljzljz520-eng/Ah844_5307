package model

import "fmt"

type Vehicle struct {
	ID       string
	Plate    string
	Capacity int
	Status   string
}
type Driver struct {
	ID      string
	Name    string
	License string
	Active  bool
}
type Route struct {
	ID          string
	Origin      string
	Destination string
	Distance    int
}
type Waybill struct {
	ID        string
	VehicleID string
	DriverID  string
	RouteID   string
	Weight    int
	Status    string
}
type Schedule struct {
	ID        string
	VehicleID string
	DriverID  string
	RouteID   string
	WaybillID string
	State     string
}
type Table struct {
	Name       string
	Columns    []string
	Types      []string
	PrimaryKey string
}
type SourceSchema struct {
	Database string
	Tables   []Table
}
type TargetDDL struct {
	SQL    string
	Tables []Table
}
type Difference struct {
	Table  string
	Kind   string
	Detail string
}
type Conversion struct {
	ID          string
	Source      SourceSchema
	Target      TargetDDL
	Differences []Difference
	State       string
	Confirmed   bool
}

func (t Table) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("table name required")
	}
	if len(t.Columns) == 0 {
		return fmt.Errorf("columns required")
	}
	if len(t.Columns) != len(t.Types) {
		return fmt.Errorf("column type mismatch")
	}
	return nil
}
func (s SourceSchema) Find(name string) (Table, bool) {
	for _, t := range s.Tables {
		if t.Name == name {
			return t, true
		}
	}
	return Table{}, false
}
func (s SourceSchema) Validate() error {
	if s.Database == "" {
		return fmt.Errorf("database required")
	}
	if len(s.Tables) == 0 {
		return fmt.Errorf("tables required")
	}
	for _, t := range s.Tables {
		if err := t.Validate(); err != nil {
			return err
		}
	}
	return nil
}
func (c Conversion) CanConfirm() bool { return c.State == "draft" && len(c.Target.Tables) > 0 }
func (c Conversion) CanExecute() bool { return c.State == "confirmed" && c.Confirmed }
func (w Waybill) IsValid() bool {
	return w.ID != "" && w.VehicleID != "" && w.DriverID != "" && w.RouteID != "" && w.Weight > 0
}
func (v Vehicle) IsAvailable() bool { return v.Status == "available" && v.Capacity > 0 }
func (d Driver) IsEligible() bool   { return d.Active && d.License != "" }
func (r Route) IsUsable() bool      { return r.Origin != "" && r.Destination != "" && r.Distance > 0 }
