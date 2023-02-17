package structx

import (
	json "github.com/json-iterator/go"
)

type Byter interface {
	Bytes() ([]byte, error)
}

// BTrans  Struct Byte to Struct
func BTrans[T any](s Byter) (res T, err error) {

	var (
		rc []byte
	)

	if rc, err = s.Bytes(); nil != err {

		return

	}

	if 0 == len(rc) {

		return

	}

	if err = json.Unmarshal(rc, &res); nil != err {

		return

	}

	return
}
