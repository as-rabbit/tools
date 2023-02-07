package structx

import (
	"encoding/json"
)

type StructToByte interface {
	ToByte() []byte
}

//  Any Type to Struct

func Trans[T comparable](s StructToByte) (res T, err error) {

	if err = json.Unmarshal(s.ToByte(), &res); nil != err {

		return

	}

	return
}
