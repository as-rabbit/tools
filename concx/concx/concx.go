package concx

import (
	"context"
	"github.com/sourcegraph/conc/pool"
)

func Tasks[T any, R any](ctx context.Context, tasks []T, f func(p T) (r R, err error), n int) (r []R, err error) {

	p := pool.NewWithResults[R]().WithContext(ctx).WithMaxGoroutines(n)

	for _, task := range tasks {

		tk := task

		p.Go(func(context.Context) (R, error) {

			return f(tk)

		})

	}

	r, err = p.Wait()

	return
}

func Streams[T any, R any](ctx context.Context) {

	return
}
