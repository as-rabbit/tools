package string

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"tool/convx/integer"
)

// ToI String To Integer
func ToI[T constraints.Integer](s string) T {

	v, _ := ToIE[T](s)

	return v
}

// ToIE String To Integer
func ToIE[T constraints.Integer](s string) (r T, err error) {

	var t T

	switch any(t).(type) {

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:

		var v int64

		v, err = strconv.ParseInt(s, 10, 64)

		r = integer.ToAnyI[T](v)

	default:

		return 0, fmt.Errorf("the type %T isn't supported", t)
	}

	return
}
