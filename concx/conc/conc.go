package conc

import (
	"context"
	"github.com/sourcegraph/conc/pool"
	"github.com/sourcegraph/conc/stream"
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

func Streams(tasks []func(), n int) {

	s := stream.New().WithMaxGoroutines(n)
	for _, task := range tasks {
		tt := task
		s.Go(func() stream.Callback {

			return tt

		})
	}

	s.Wait()

	return
}
