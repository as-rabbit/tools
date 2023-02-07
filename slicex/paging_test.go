package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPagination(t *testing.T) {

	d := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tests := []struct {
		Case           string
		Data           []int64
		Page           int
		Size           int
		TestIndexStart int
		TestIndexEnd   int
		TestTotal      int
	}{
		{
			Case:           "next page",
			Data:           d,
			Page:           1,
			Size:           5,
			TestIndexStart: 0,
			TestIndexEnd:   5,
			TestTotal:      2,
		},
		{
			Case:           "overstep page",
			Data:           d,
			Page:           3,
			Size:           5,
			TestIndexStart: 0,
			TestIndexEnd:   0,
			TestTotal:      2,
		},
		{
			Case:           "page is zero",
			Data:           d,
			Page:           0,
			Size:           5,
			TestIndexStart: 0,
			TestIndexEnd:   0,
			TestTotal:      1,
		},
	}

	for _, test := range tests {
		t.Run(test.Case, func(t *testing.T) {

			resp1 := Pagination[int, int64](test.Page, test.Size, d)

			assert.Equal(t, test.TestIndexStart, resp1.IndexStart)
			assert.Equal(t, test.TestIndexEnd, resp1.IndexEnd)
			assert.Equal(t, test.TestTotal, resp1.PageTotal)
			assert.Equal(t, d[test.TestIndexStart:test.TestIndexEnd], resp1.Data)

		})
	}

}

func TestTop(t *testing.T) {

	d := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, d[0:5], Top[int, int64](5, d))

}
