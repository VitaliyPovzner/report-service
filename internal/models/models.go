package models

import (
	"fmt"
	"time"
)

type FlexibleDateTime time.Time

type Filter struct {
	Operand  string `json:"operand"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

const (
	dateFormat     = "2006-01-02"
	dateTimeFormat = "2006-01-02T15:04:05"
)

func (d *FlexibleDateTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	t, err := time.Parse(dateTimeFormat, str)
	if err == nil {
		*d = FlexibleDateTime(t)
		return nil
	}

	t, err = time.Parse(dateFormat, str)
	if err == nil {
		*d = FlexibleDateTime(t)
		return nil
	}

	return fmt.Errorf("could not parse date or datetime: %v", err)
}
func (d FlexibleDateTime) Format(layout string) string {
	return time.Time(d).Format(layout)
}

type AggregationRequest struct {
	Dimensions []string         `json:"dimensions"`
	Metrics    []string         `json:"metrics"`
	DateFrom   FlexibleDateTime `json:"dateFrom"`
	DateTo     FlexibleDateTime `json:"dateTo"`
	Filters    []Filter         `json:"filters"`
	Breakdown  string           `json:"breakdown,omitempty"`
}
