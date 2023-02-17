package errorx

import (
	"errors"
)

var FilterErr = errors.New("data is abnormal")

func IsFilter(err error) bool {

	return errors.Is(err, FilterErr)

}
