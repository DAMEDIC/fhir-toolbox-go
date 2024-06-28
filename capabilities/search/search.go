package search

import (
	"fhir-toolbox/model"
	"fmt"
	"time"
)

type Cursor string

type Result struct {
	Resources []model.Resource
	Next      Cursor
}

type Capabilities struct {
	Parameters map[string]Desc
	Includes   []string
}

type Desc struct {
	Type Type
}

type Type string

const (
	Number    Type = "number"
	Date      Type = "date"
	String    Type = "string"
	Token     Type = "token"
	Reference Type = "reference"
	Composite Type = "composite"
	Quantity  Type = "quantity"
	Uri       Type = "uri"
	Special   Type = "special"
)

type Options struct {
	Parameters Parameters
	Includes   []string
	Count      int
	Cursor     Cursor
}

type Parameters = map[string]AndList
type AndList = []OrList
type OrList = []Value

type Value struct {
	Prefix        Prefix
	DatePrecision DatePrecision
	Value         any
}

type Prefix = string

const (
	Equal          Prefix = "eq"
	NotEqual       Prefix = "ne"
	GreaterThan    Prefix = "gt"
	LessThan       Prefix = "lt"
	GreaterOrEqual Prefix = "ge"
	LessOrEqual    Prefix = "le"
)

var KnownPrefixes = []Prefix{Equal, NotEqual, GreaterThan, LessThan, GreaterOrEqual, LessOrEqual}

type DatePrecision string

const (
	YearPrecision       DatePrecision = "year"
	MonthPrecision      DatePrecision = "month"
	DayPrecision        DatePrecision = "day"
	HourMinutePrecision DatePrecision = "hourMinute"
	FullTimePrecision   DatePrecision = "time"
)

const (
	OnlyYearFormat   = "2006"
	UpToMonthFormat  = "2006-01"
	UpToDayFormat    = "2006-01-02"
	HourMinuteFormat = "2006-01-02T15:04Z07:00"
	FullTimeFormat   = "2006-01-02T15:04:05.999999999Z07:00"
)

func (v Value) String() string {
	var s string

	switch v.Value.(type) {
	case time.Time:
		date, isDate := v.Value.(time.Time)
		if isDate {
			s = formatDate(date, v.DatePrecision)
		}
	default:
		s = fmt.Sprintf("%v", v.Value)
	}

	return v.Prefix + s
}

func formatDate(v time.Time, p DatePrecision) string {
	switch p {
	case YearPrecision:
		return v.Format(OnlyYearFormat)
	case MonthPrecision:
		return v.Format(UpToMonthFormat)
	case DayPrecision:
		return v.Format(UpToDayFormat)
	case HourMinutePrecision:
		return v.Format(HourMinuteFormat)
	default:
		return v.Format(FullTimeFormat)
	}
}
