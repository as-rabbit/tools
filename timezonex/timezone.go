package timezonex

import (
	"time"
	"tool/config/timezone"
)

func CountryTrans(country string, t time.Time) (res time.Time, err error) {

	var (
		countryTimeZone string
		timeZone        *time.Location
	)

	countryTimeZone = timezone.CountryTimezone(country)

	if timeZone, err = time.LoadLocation(countryTimeZone); nil != err {

		return time.Time{}, err

	}

	return transTime(timeZone, t), nil
}

func CountryRangeTrans(country string, startTime time.Time, endTime time.Time) (s time.Time, e time.Time, err error) {

	var (
		countryTimeZone string
		timeZone        *time.Location
	)

	countryTimeZone = timezone.CountryTimezone(country)

	if timeZone, err = time.LoadLocation(countryTimeZone); nil != err {

		return time.Time{}, time.Time{}, err

	}

	return transTime(timeZone, startTime), transTime(timeZone, endTime), nil

}

// transTime Timezone Time Trans To Location Time
func transTime(tz *time.Location, t time.Time) time.Time {

	return t.In(tz)

}
