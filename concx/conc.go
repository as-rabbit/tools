package concx

import (
	"context"
	"errors"
	cp "github.com/sourcegraph/conc/pool"
	"gopkg.in/go-playground/pool.v3"
)

func PoolBatchTask[T any, R any](p pool.Pool, tasks []T, f func(p any) pool.WorkFunc) (r []R, err error) {

	r = make([]R, len(tasks))
	i := 0

	batch := p.Batch()

	for _, taskId := range tasks {

		batch.Queue(f(taskId))

	}

	batch.QueueComplete()

	for resp := range batch.Results() {

		if err = resp.Error(); nil != err {

			return []R{}, err

		}

		tmpRes, ok := resp.Value().(R)

		if !ok {

			return []R{}, errors.New("task response type err")

		}

		r[i] = tmpRes

		i++
	}

	return
}

func ConBatchTask[T any, R any](ctx context.Context, tasks []T, f func(p T) (r R, err error), n int) (r []R, err error) {

	p := cp.NewWithResults[R]().WithContext(ctx).WithMaxGoroutines(n)

	for _, taskId := range tasks {

		p.Go(func(context.Context) (R, error) {

			return f(taskId)

		})

	}

	r, err = p.Wait()

	return
}
