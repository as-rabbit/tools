package integer

import (
	"golang.org/x/exp/constraints"
)

// ToAnyI int64 to integer
func ToAnyI[T constraints.Integer](i int64) (r T) {

	return T(i)

}
