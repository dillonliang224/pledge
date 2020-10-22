package moment

import (
	"time"
)

// 实现类似于js里moment功能的库

type Moment struct {
	time.Time
}

func (moment *Moment) AddTime(num int, s string) time.Time {
	switch s {
	case "year":
		return moment.AddDate(num, 0, 0)
	case "month":
		return moment.AddDate(0, num, 0)
	case "day":
		return moment.AddDate(0, 0, num)
	case "hour":
		return moment.Add(time.Hour * time.Duration(num))
	default:
		return time.Now()
	}
}

func (moment *Moment) SubTime(num int, s string) time.Time {
	switch s {
	case "year":
		return moment.AddDate(-num, 0, 0)
	case "month":
		return moment.AddDate(0, -num, 0)
	case "day":
		return moment.AddDate(0, 0, -num)
	case "hour":
		return moment.Add(time.Hour * time.Duration(-num))
	default:
		return time.Now()
	}
}

func (moment *Moment) StartOf(s string) time.Time {
	switch s {
	case "year":
		y, _, _ := moment.Date()
		return time.Date(y, time.January, 1, 0, 0, 0, 0, moment.Location())
	case "month":
		y, m, _ := moment.Date()
		return time.Date(y, m, 1, 0, 0, 0, 0, moment.Location())
	case "day":
		y, m, d := moment.Date()
		return time.Date(y, m, d, 0, 0, 0, 0, moment.Location())
	case "hour":
		y, m, d := moment.Date()
		return time.Date(y, m, d, moment.Time.Hour(), 0, 0, 0, moment.Location())
	default:
		return time.Now()
	}
}

func (moment *Moment) EndOf(s string) time.Time {
	switch s {
	case "year":
		return moment.StartOf("year").AddDate(1, 0, 0).Add(-time.Nanosecond)
	case "month":
		return moment.StartOf("month").AddDate(0, 1, 0).Add(-time.Nanosecond)
	case "day":
		y, m, d := moment.Date()
		return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), moment.Location())
	case "hour":
		return moment.StartOf("hour").Add(time.Hour - time.Nanosecond)
	default:
		return time.Now()
	}
}

func (moment *Moment) IsBefore(t time.Time) bool {
	return moment.Before(t)
}

func (moment *Moment) IsAfter(t time.Time) bool {
	return moment.After(t)
}

func (moment *Moment) IsBetween(b time.Time, a time.Time) bool {
	return moment.Before(b) && moment.After(a)
}
