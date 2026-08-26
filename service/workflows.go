package service

import (
	"fmt"
	"logisticsconvert/differ"
	"logisticsconvert/model"
)

func (s *Service) GenerateDDL(id string) (model.TargetDDL, error) {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return model.TargetDDL{}, e
	}
	if len(c.Source.Tables) == 0 {
		return model.TargetDDL{}, fmt.Errorf("no tables")
	}
	return differ.BuildDDL(c.Source), nil
}
func (s *Service) BuildPlan(id string) ([]differ.PlanStep, error) {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return nil, e
	}
	return differ.Plan(c.Source, targetSource(c)), nil
}
func (s *Service) Validate(id string) error {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return e
	}
	if len(c.Differences) > 0 && differ.HasBlocking(c.Differences) {
		return fmt.Errorf("blocking differences")
	}
	return nil
}
func (s *Service) Duplicate(id, newID string) error {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return e
	}
	if newID == "" {
		return fmt.Errorf("new id required")
	}
	c.ID = newID
	c.State = "draft"
	c.Confirmed = false
	return s.Store.PutConversion(c)
}
func (s *Service) Reset(id string) error { return s.SafeTransition(id, "draft") }
func (s *Service) ExecutePlan(id string, confirm bool) (string, error) {
	if !confirm {
		return "", fmt.Errorf("confirmation required")
	}
	if _, e := s.Execute(id, true); e != nil {
		return "", e
	}
	ddl, e := s.GenerateDDL(id)
	if e != nil {
		return "", e
	}
	return ddl.SQL, nil
}
