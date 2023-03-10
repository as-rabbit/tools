package conc

import (
	"context"
	"github.com/stretchr/testify/assert"

	"testing"
)

var tasks = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestTasks(t *testing.T) {

	ctx := context.TODO()

	res, err := Tasks[int64, int64](ctx, tasks, func(p int64) (r int64, err error) {

		return p, nil

	}, 10)

	assert.Equal(t, len(res), len(tasks))
	assert.Equal(t, nil, err)
}
