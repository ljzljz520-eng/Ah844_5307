package store

import (
	"encoding/json"
	"logisticsconvert/model"
	"os"
)

func (s *Store) Export(path string) error {
	items, e := s.ListConversions()
	if e != nil {
		return e
	}
	b, e := json.MarshalIndent(items, "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(path, b, 0600)
}
func (s *Store) Import(path string) error {
	b, e := os.ReadFile(path)
	if e != nil {
		return e
	}
	var items []model.Conversion
	if e = json.Unmarshal(b, &items); e != nil {
		return e
	}
	for _, c := range items {
		if e = s.PutConversion(c); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) Exists(id string) bool { _, e := s.GetConversion(id); return e == nil }
func (s *Store) Replace(c model.Conversion) error {
	if c.ID == "" {
		return os.ErrInvalid
	}
	return s.PutConversion(c)
}
func (s *Store) Count() (int, error) { items, e := s.ListConversions(); return len(items), e }
