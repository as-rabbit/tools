package structx

import (
	json "github.com/json-iterator/go"
)

type StructToByte interface {
	Bytes() []byte
}

// BTrans  Struct to Struct
func BTrans[T any](s StructToByte) (res T, err error) {

	if err = json.Unmarshal(s.Bytes(), &res); nil != err {

		return

	}

	return
}
