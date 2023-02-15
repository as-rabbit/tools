package timezonex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tool/config/timezone"
)

func TestCountryTrans(t *testing.T) {
	var (
		err              error
		localTime        time.Time
		thTimeZone       time.Time
		expectedTimeZone time.Time
		timeZone         *time.Location
		timeCountry      = "TH"
	)

	countryTimeZone := timezone.CountryTimezone(timeCountry)

	timeZone, err = time.LoadLocation(countryTimeZone)

	assert.Equal(t, nil, err)

	thTimeZone, err = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", timeZone)

	assert.Equal(t, nil, err)

	localTime, err = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 01:00:00", time.Local)

	assert.Equal(t, nil, err)

	expectedTimeZone, err = CountryTrans(timeCountry, localTime)

	assert.Equal(t, expectedTimeZone, thTimeZone)

}

func TestCountryRangeTrans(t *testing.T) {
	var (
		stratTime       time.Time
		endTime         time.Time
		stratThTimeZone time.Time
		endThTimeZone   time.Time
		timeZone        *time.Location
		timeCountry     = "TH"
	)
	stratTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 01:00:00", time.Local)
	endTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 01:00:00", time.Local)

	countryTimeZone := timezone.CountryTimezone(timeCountry)
	timeZone, _ = time.LoadLocation(countryTimeZone)
	stratThTimeZone, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", timeZone)
	endThTimeZone, _ = time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", timeZone)

	type args struct {
		country   string
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantS   time.Time
		wantE   time.Time
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "timezone:th",
			args: args{
				country:   timeCountry,
				startTime: stratTime,
				endTime:   endTime,
			},
			wantS: stratThTimeZone,
			wantE: endThTimeZone,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {

				return nil == err

			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotE, err := CountryRangeTrans(tt.args.country, tt.args.startTime, tt.args.endTime)
			if !tt.wantErr(t, err, fmt.Sprintf("CountryRangeTrans(%v, %v, %v)", tt.args.country, tt.args.startTime, tt.args.endTime)) {
				return
			}
			assert.Equalf(t, tt.wantS, gotS, "CountryRangeTrans(%v, %v, %v)", tt.args.country, tt.args.startTime, tt.args.endTime)
			assert.Equalf(t, tt.wantE, gotE, "CountryRangeTrans(%v, %v, %v)", tt.args.country, tt.args.startTime, tt.args.endTime)
		})
	}
}
