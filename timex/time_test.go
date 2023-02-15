package timex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDayZeroTime(t *testing.T) {
	type args struct {
		d time.Time
	}

	xt := time.Now()
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "NextZeroTime",
			args: args{
				d: xt,
			},
			want: time.Date(xt.Year(), xt.Month(), xt.Day(), 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DayZeroTime(tt.args.d), "DayZeroTime(%v)", tt.args.d)
		})
	}
}

func TestDayNextZeroTime(t *testing.T) {
	type args struct {
		d time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "NextZeroTime",
			args: args{
				d: time.Date(2023, 01, 01, 23, 59, 59, 0, time.Local),
			},
			want: time.Date(2023, 01, 02, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DayNextZeroTime(tt.args.d), "DayNextZeroTime(%v)", tt.args.d)
		})
	}
}

func TestDayLastTime(t *testing.T) {
	type args struct {
		d time.Time
	}

	xt := time.Now()

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "DayLastTime",
			args: args{
				d: time.Now(),
			},
			want: time.Date(xt.Year(), xt.Month(), xt.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DayLastTime(tt.args.d), "DayLastTime(%v)", tt.args.d)
		})
	}
}
