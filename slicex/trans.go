package slicex

import (
	json "github.com/json-iterator/go"
)

// Slice String To Transform Any Type

func STrans[T any, R any](collection []T) (res []R, err error) {
	var (
		rc []byte
	)

	if rc, err = json.Marshal(collection); nil != err {

		return

	}

	if err = json.Unmarshal(rc, &res); nil != err {

		return []R{}, err

	}

	return
}
