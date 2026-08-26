package service

import (
	"fmt"
	"logisticsconvert/model"
)

func ValidateWaybill(w model.Waybill, v model.Vehicle, d model.Driver, r model.Route) error {
	if !w.IsValid() {
		return fmt.Errorf("invalid waybill")
	}
	if !v.IsAvailable() {
		return fmt.Errorf("vehicle unavailable")
	}
	if !d.IsEligible() {
		return fmt.Errorf("driver ineligible")
	}
	if !r.IsUsable() {
		return fmt.Errorf("route unusable")
	}
	if w.Weight > v.Capacity {
		return fmt.Errorf("over capacity")
	}
	return nil
}
func RouteDistance(r model.Route) int {
	if r.Distance < 0 {
		return 0
	}
	return r.Distance
}
func ScheduleReady(sc model.Schedule) bool { return sc.ID != "" && sc.State == "planned" }
func AdvanceSchedule(sc *model.Schedule, next string) error {
	allowed := map[string]string{"planned": "dispatched", "dispatched": "delivered", "delivered": "closed"}
	want, ok := allowed[sc.State]
	if !ok || want != next {
		return fmt.Errorf("invalid schedule transition")
	}
	sc.State = next
	return nil
}
func ValidateSchedule(sc model.Schedule, w model.Waybill) error {
	if !ScheduleReady(sc) {
		return fmt.Errorf("schedule not planned")
	}
	if w.Status == "cancelled" {
		return fmt.Errorf("cancelled waybill")
	}
	return nil
}
