package string

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestToIE(t *testing.T) {

	a, e := ToIE[int64]("1")

	assert.Equal(t, int64(1), a, e)

}

func BenchmarkStringToInt(b *testing.B) {

	for i := 0; i < b.N; i++ {

		strconv.ParseInt("1", 10, 64)

	}

}

func BenchmarkToI(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_ = ToI[int64]("1")

	}

}

func BenchmarkToIE(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_, _ = ToIE[int64]("1")

	}

}
