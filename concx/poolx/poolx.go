package poolx

import (
	"context"
	"errors"
	"gopkg.in/go-playground/pool.v3"
	"tool/errorx"
)

func work[T any](ctx context.Context, req T, f func(ctx context.Context, req T) (interface{}, error)) pool.WorkFunc {

	return func(wu pool.WorkUnit) (interface{}, error) {

		return f(ctx, req)

	}
}

type WorkUint[P any, R any] struct {
	wu pool.WorkUnit
}

func NewWorkUnit[T any, R any](ctx context.Context, p pool.Pool, req T, f func(ctx context.Context, req T) (interface{}, error)) *WorkUint[T, R] {

	return &WorkUint[T, R]{
		wu: p.(pool.Pool).Queue(work(ctx, req, f)),
	}

}

func (w *WorkUint[T, R]) Response() (res R, err error) {

	if res, err = dealResp[R](w.wu); nil != err {

		return
	}

	return res, nil
}

// WorkUnit Resp Value
func dealResp[R any](w pool.WorkUnit) (r R, err error) {

	w.Wait()

	if err = w.Error(); nil != err {

		return r, err

	}

	r, ok := w.Value().(R)

	if !ok {

		return r, errors.New("task response type err")

	}

	return
}

func Steams[T any, R any](ctx context.Context, p pool.Pool, steams []T, f func(ctx context.Context, req T) (interface{}, error)) (r []R, err error) {

	r, x := make([]R, len(steams)), make([]pool.WorkUnit, len(steams))

	for taskId, task := range steams {

		tk := task

		x[taskId] = p.Queue(work(ctx, tk, f))

	}

	for i, rsp := range x {

		var (
			tr R
			e  error
		)

		if tr, e = dealResp[R](rsp); errorx.IsFilter(e) {

			continue

		}

		if nil != e {

			return []R{}, e

		}

		r[i] = tr
	}

	return
}

func Tasks[PK comparable, PV any, R any](ctx context.Context, p pool.Pool, tasks map[PK]PV, f func(ctx context.Context, req PV) (interface{}, error)) (r map[PK]R, err error) {

	r, x := make(map[PK]R, len(tasks)), make(map[PK]pool.WorkUnit, len(tasks))

	for taskId, task := range tasks {

		tk := task

		x[taskId] = p.Queue(work(ctx, tk, f))

	}

	for i, rsp := range x {

		var (
			tr R
			e  error
		)

		if tr, e = dealResp[R](rsp); errorx.IsFilter(e) {

			continue

		}

		if nil != e {

			return map[PK]R{}, e

		}

		r[i] = tr
	}

	return
}
