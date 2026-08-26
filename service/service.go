package service

import (
	"fmt"
	"logisticsconvert/differ"
	"logisticsconvert/model"
	"logisticsconvert/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Create(id string, src model.SourceSchema, names []string) (model.Conversion, error) {
	if id == "" {
		return model.Conversion{}, fmt.Errorf("id required")
	}
	selected := model.SourceSchema{Database: src.Database}
	for _, n := range names {
		if t, ok := src.Find(n); ok {
			selected.Tables = append(selected.Tables, t)
		}
	}
	if err := selected.Validate(); err != nil {
		return model.Conversion{}, err
	}
	c := model.Conversion{ID: id, Source: selected, Target: differ.BuildDDL(selected), State: "draft"}
	c.Differences = differ.Compare(selected, targetSource(c))
	if err := s.Store.PutConversion(c); err != nil {
		return c, err
	}
	return c, nil
}
func targetSource(c model.Conversion) model.SourceSchema {
	return model.SourceSchema{Database: "target", Tables: c.Target.Tables}
}
func (s *Service) Confirm(id string) (model.Conversion, error) {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return c, e
	}
	if !c.CanConfirm() {
		return c, fmt.Errorf("cannot confirm in state %s", c.State)
	}
	c.State = "confirmed"
	c.Confirmed = true
	if e = s.Store.PutConversion(c); e != nil {
		return c, e
	}
	return c, nil
}
func (s *Service) Execute(id string, confirmed bool) (model.Conversion, error) {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return c, e
	}
	if !confirmed || !c.CanExecute() {
		return c, fmt.Errorf("second confirmation required")
	}
	c.State = "executed"
	return c, s.Store.PutConversion(c)
}
func (s *Service) Transition(id, next string) error {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return e
	}
	c.State = next
	return s.Store.PutConversion(c)
}
func (s *Service) SafeTransition(id, next string) error {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return e
	}
	allowed := map[string][]string{"draft": {"confirmed"}, "confirmed": {"executed"}, "executed": {}}
	ok := false
	for _, v := range allowed[c.State] {
		if v == next {
			ok = true
		}
	}
	if !ok {
		return fmt.Errorf("invalid transition %s to %s", c.State, next)
	}
	c.State = next
	return s.Store.PutConversion(c)
}
func (s *Service) Report(id string) (string, error) {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return "", e
	}
	return differ.Report(c.Differences), nil
}
func (s *Service) Delete(id string) error { return s.Store.DeleteConversion(id) }
