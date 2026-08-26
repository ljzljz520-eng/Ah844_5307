package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"logisticsconvert/model"
	"os"
)

var bucket = []byte("logistics")

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	err = db.Update(func(tx *bbolt.Tx) error { _, e := tx.CreateBucketIfNotExists(bucket); return e })
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) Path() string { return s.path }
func (s *Store) PutConversion(c model.Conversion) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Put([]byte(c.ID), data) })
}
func (s *Store) GetConversion(id string) (model.Conversion, error) {
	var c model.Conversion
	err := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(bucket).Get([]byte(id))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, &c)
	})
	return c, err
}
func (s *Store) ListConversions() ([]model.Conversion, error) {
	var out []model.Conversion
	err := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(_, v []byte) error {
			var c model.Conversion
			if err := json.Unmarshal(v, &c); err != nil {
				return err
			}
			out = append(out, c)
			return nil
		})
	})
	return out, err
}
func (s *Store) DeleteConversion(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(bucket).Delete([]byte(id)) })
}
func (s *Store) MustGet(id string) model.Conversion {
	c, e := s.GetConversion(id)
	if e != nil {
		panic(fmt.Sprintf("conversion %s: %v", id, e))
	}
	return c
}
