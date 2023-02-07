package timex

import (
	"time"
)

// DayZeroTime Get Time.time Day 00:00:00
func DayZeroTime(d time.Time) time.Time {

	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())

}

// DayLastTime Get Time.time Day 23:59:59.999999999
func DayLastTime(d time.Time) time.Time {

	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), d.Location())

}

// DayNextZeroTime Get Time.time Day +1 00:00:00
func DayNextZeroTime(d time.Time) time.Time {

	return DayZeroTime(d).AddDate(0, 0, 1)

}

// MonthZeroTime Get Time.time Month 00:00:00
func MonthZeroTime(d time.Time) time.Time {

	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())

}

// MonthLastTime Get Time.time Month 23:59:59.999999999
func MonthLastTime(d time.Time) time.Time {

	return MonthZeroTime(d).AddDate(0, 1, 0).Add(-time.Nanosecond)

}

// MonthNextZeroTime Get Time.time Month +1 00:00:00
func MonthNextZeroTime(d time.Time) time.Time {

	return MonthZeroTime(d).AddDate(0, 1, 0)

}
