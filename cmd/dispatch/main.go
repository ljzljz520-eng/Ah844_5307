package main

import (
	"encoding/json"
	"fmt"
	"logisticsconvert/model"
	"logisticsconvert/service"
	"logisticsconvert/store"
	"os"
)

func main() {
	path := "dispatch.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	s, e := store.Open(path)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer s.Close()
	src := model.SourceSchema{Database: "mysql", Tables: []model.Table{{Name: "vehicles", Columns: []string{"id", "plate"}, Types: []string{"INT", "VARCHAR"}}, {Name: "drivers", Columns: []string{"id", "name"}, Types: []string{"INT", "VARCHAR"}}, {Name: "routes", Columns: []string{"id", "origin"}, Types: []string{"INT", "VARCHAR"}}, {Name: "waybills", Columns: []string{"id", "weight"}, Types: []string{"INT", "INT"}}}}
	svc := service.New(s)
	c, e := svc.Create("demo", src, []string{"vehicles", "drivers", "routes", "waybills"})
	if e == nil {
		c, _ = svc.Confirm(c.ID)
		c, e = svc.Execute(c.ID, true)
	}
	b, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(b))
	if e != nil {
		fmt.Println(e)
	}
}
