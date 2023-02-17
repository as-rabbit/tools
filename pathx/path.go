package pathx

import (
	"fmt"
	"os"
)

// CreateIfNotExist creates a file if it is not exists.
func CreateIfNotExist(file string) (*os.File, error) {
	_, err := os.Stat(file)

	if !os.IsNotExist(err) {

		return nil, fmt.Errorf("%s already exist", file)

	}

	return os.Create(file)
}

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {

	_, err := os.Stat(file)

	return nil == err
}

// IsDir returns true If the specified directory exists.
func IsDir(path string) bool {
	var (
		err error
		s   os.FileInfo
	)

	if s, err = os.Stat(path); nil != err {

		return false
	}

	return s.IsDir()
}
