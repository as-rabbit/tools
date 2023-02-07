package integer

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

// ToAnyI int64 to integer
func ToAnyI[T constraints.Integer](i int64) (r T) {

	return T(i)

}

// ToAnyS Integer to string
func ToAnyS[T constraints.Integer](i T) (r string) {

	return strconv.FormatInt(int64(i), 10)

}
