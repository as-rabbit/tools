package integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAnyI(t *testing.T) {

	assert.Equal(t, int32(1), ToAnyI[int32](1))

}
