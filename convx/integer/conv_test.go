package integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAnyI(t *testing.T) {

	assert.Equal(t, 1, ToAnyI[int](1))
	assert.Equal(t, int8(1), ToAnyI[int8](1))
	assert.Equal(t, int16(1), ToAnyI[int16](1))
	assert.Equal(t, int32(1), ToAnyI[int32](1))
	assert.Equal(t, int64(1), ToAnyI[int64](1))

	assert.Equal(t, uint(1), ToAnyI[uint](1))
	assert.Equal(t, uint8(1), ToAnyI[uint8](1))
	assert.Equal(t, uint16(1), ToAnyI[uint16](1))
	assert.Equal(t, uint32(1), ToAnyI[uint32](1))
	assert.Equal(t, uint64(1), ToAnyI[uint64](1))

}
