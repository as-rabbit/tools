package concx

import (
	"context"
	"gopkg.in/go-playground/pool.v3"
	"testing"
	"tool/concx/conc"
	"tool/concx/poolx"
)

var tasks = map[int64]int64{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10}
var steams = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func BenchmarkPoolXTasks(b *testing.B) {

	ctx := context.TODO()
	p := pool.NewLimited(10)

	defer p.Close()

	for i := 0; i < b.N; i++ {

		_, _ = poolx.Tasks[int64, int64, int64](ctx, p, tasks, func(ctx context.Context, req int64) (interface{}, error) {

			return req, nil

		})

	}
}

func BenchmarkConcTasks(b *testing.B) {

	ctx := context.TODO()

	for i := 0; i < b.N; i++ {

		_, _ = conc.Tasks[int64, int64](ctx, steams, func(p int64) (r int64, err error) {

			return p, nil

		}, 10)

	}

}

func BenchmarkConcStreams(b *testing.B) {

	for i := 0; i < b.N; i++ {

		conc.Streams([]func(){}, 10)

	}

}
