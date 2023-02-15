package pathx

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateIfNotExist(t *testing.T) {

	fmt.Println(os.ExpandEnv("$HOME/golang-server.log"))
	//err := CreateIfNotExist("./test/test.go")
	//fmt.Println(err)
}
