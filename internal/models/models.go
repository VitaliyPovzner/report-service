package models

import (
	"time"
	"fmt"
)

type FlexibleDateTime time.Time

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

type AggregationRequest struct {
	Dimensions []string         `json:"dimensions"`
	Metrics    []string         `json:"metrics"`
	DateFrom   FlexibleDateTime `json:"dateFrom"`
	DateTo     FlexibleDateTime `json:"dateTo"`
	Breakdown  string           `json:"breakdown,omitempty"`
}