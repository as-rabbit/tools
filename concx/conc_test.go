package concx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/pool.v3"
	"testing"
)

var tasks = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestPoolBatchTask(t *testing.T) {

	p := pool.NewLimited(10)
	defer p.Close()

	res, err := PoolBatchTask[int64, int64](p, tasks, func(params any) pool.WorkFunc {

		return func(wu pool.WorkUnit) (interface{}, error) {

			return params, nil

		}
	})

	assert.Equal(t, len(res), len(tasks))
	assert.Equal(t, nil, err)
}

func BenchmarkPoolBatchTask(b *testing.B) {

	p := pool.NewLimited(10)

	defer p.Close()

	for i := 0; i < b.N; i++ {

		_, _ = PoolBatchTask[int64, int64](p, tasks, func(params any) pool.WorkFunc {

			return func(wu pool.WorkUnit) (interface{}, error) {

				return params, nil

			}
		})

	}
}

func TestConBatchTask(t *testing.T) {

	ctx := context.TODO()

	res, err := ConBatchTask[int64, int64](ctx, tasks, func(p int64) (r int64, err error) {

		return p, nil

	}, 10)

	assert.Equal(t, len(res), len(tasks))
	assert.Equal(t, nil, err)
}

func BenchmarkConBatchTask(b *testing.B) {

	ctx := context.TODO()

	for i := 0; i < b.N; i++ {

		_, _ = ConBatchTask[int64, int64](ctx, tasks, func(p int64) (r int64, err error) {

			return p, nil

		}, 10)

	}

}
