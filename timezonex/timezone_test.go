package timezonex

import (
	"fmt"
	"testing"
	"time"
)

func TestCountryRangeTrans(t *testing.T) {

	tt := time.Now()

	tt, err := CountryTrans("TH", tt)

	fmt.Println(tt, err)
}
