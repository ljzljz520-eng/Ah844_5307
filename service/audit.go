package service

import (
	"fmt"
	"logisticsconvert/model"
	"strings"
)

type AuditEvent struct {
	ConversionID string
	Action       string
	From         string
	To           string
}

func DescribeTransition(c model.Conversion, next string) AuditEvent {
	return AuditEvent{ConversionID: c.ID, Action: "transition", From: c.State, To: next}
}
func FormatAudit(e AuditEvent) string {
	return fmt.Sprintf("%s %s %s->%s", e.ConversionID, e.Action, e.From, e.To)
}
func AllowedTransitions(state string) []string {
	switch state {
	case "draft":
		return []string{"confirmed"}
	case "confirmed":
		return []string{"executed"}
	case "executed":
		return []string{}
	default:
		return nil
	}
}
func IsAllowed(state, next string) bool {
	for _, v := range AllowedTransitions(state) {
		if v == next {
			return true
		}
	}
	return false
}
func (s *Service) Audit(id string) string {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return ""
	}
	return strings.Join([]string{c.ID, c.State, fmt.Sprint(len(c.Differences))}, "|")
}
func (s *Service) TransitionWithAudit(id, next string) (AuditEvent, error) {
	c, e := s.Store.GetConversion(id)
	if e != nil {
		return AuditEvent{}, e
	}
	ev := DescribeTransition(c, next)
	if e = s.SafeTransition(id, next); e != nil {
		return ev, e
	}
	return ev, nil
}
func ConversionEntityNames() []string {
	return []string{"Vehicle", "Driver", "Route", "Waybill", "Schedule", "Table", "Conversion"}
}
