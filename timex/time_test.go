package timex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDayZeroTime(t *testing.T) {

	tt, _ := time.Parse("2006-01-02 15:04:05", "2006-01-01 00:00:00")

	condition, _ := time.Parse("2006-01-02 15:04:05", "2006-01-01 00:00:00")

	assert.Equal(t, tt, condition)

}

func TestDayLastTime(t *testing.T) {

	tt, _ := time.Parse("2006-01-02 15:04:05.999999999", "2006-01-01 23:59:59.999999999")

	ymd, _ := time.Parse("2006-01-02 15:04:05", "2006-01-01 00:00:00")

	assert.Equal(t, tt, DayLastTime(ymd))

}
