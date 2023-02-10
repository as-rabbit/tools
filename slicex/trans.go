package slicex

import (
	json "github.com/json-iterator/go"
)

// Slice String To Transform Any Type

func STrans[T any](collection []string) (res []T, err error) {

	res = make([]T, len(collection))

	for i, v := range collection {

		var (
			tmp = new(T)
		)

		if err = json.Unmarshal([]byte(v), &tmp); nil != err {

			return []T{}, err

		}

		res[i] = *tmp

	}

	return
}
