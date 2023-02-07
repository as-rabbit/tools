package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPagination(t *testing.T) {
	// case 1
	d := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resp1 := Pagination[int, int64](1, 5, d)

	assert.Equal(t, 0, resp1.IndexStart)
	assert.Equal(t, 5, resp1.IndexEnd)
	assert.Equal(t, 2, resp1.PageTotal)
	assert.Equal(t, d[0:5], resp1.Data)

	// case 2
	resp2 := Pagination[int, int64](3, 5, d)

	assert.Equal(t, 0, resp2.IndexStart)
	assert.Equal(t, 0, resp2.IndexEnd)
	assert.Equal(t, 2, resp2.PageTotal)
	assert.Equal(t, 0, len(resp2.Data))

	// case 3
	resp3 := Pagination[int, int64](0, 5, d)

	assert.Equal(t, 0, resp3.IndexStart)
	assert.Equal(t, 0, resp3.IndexEnd, 0)
	assert.Equal(t, 1, resp3.PageTotal)
	assert.Equal(t, 0, len(resp3.Data))
}

func TestTop(t *testing.T) {

	d := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, d[0:5], Top[int, int64](5, d))

}
