package structx

import (
	json "github.com/json-iterator/go"
)

type StructToByte interface {
	ToByte() []byte
}

// BTrans  Struct to Struct
func BTrans[T comparable](s StructToByte) (res T, err error) {

	if err = json.Unmarshal(s.ToByte(), &res); nil != err {

		return

	}

	return
}
