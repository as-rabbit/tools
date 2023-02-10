package poolx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/pool.v3"
	"testing"
)

var steams = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var tasks = map[int64]int64{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10}

type Test struct {
	Title string `json:"title"`
}

func TestWorkUint(t *testing.T) {
	var (
		p1 int64 = 1
		p2       = "2"
		p3       = 1.1
		tt       = Test{
			Title: "测试",
		}
	)

	ctx := context.TODO()
	p := pool.NewLimited(10)
	defer p.Close()

	a := NewWorkUnit[int64, int64](ctx, p, p1, func(ctx context.Context, req int64) (interface{}, error) {

		return req, nil

	})

	b := NewWorkUnit[string, string](ctx, p, p2, func(ctx context.Context, req string) (interface{}, error) {

		return req, nil

	})

	c := NewWorkUnit[float64, Test](ctx, p, p3, func(ctx context.Context, req float64) (interface{}, error) {

		return tt, nil

	})

	ar, _ := a.Response()
	br, _ := b.Response()
	cr, _ := c.Response()

	assert.Equal(t, p1, ar)
	assert.Equal(t, p2, br)
	assert.Equal(t, cr, tt)
}

func BenchmarkWorkUint(b *testing.B) {

	ctx := context.TODO()
	p := pool.NewLimited(10)

	defer p.Close()

	for i := 0; i < b.N; i++ {

		a := NewWorkUnit[int64, int64](ctx, p, 1, func(ctx context.Context, req int64) (interface{}, error) {

			return req, nil

		})
		_, _ = a.Response()
	}
}

func TestSteams(t *testing.T) {

	ctx := context.TODO()
	p := pool.NewLimited(10)
	defer p.Close()

	res, err := Steams[int64, int64](ctx, p, steams, func(ctx context.Context, req int64) (interface{}, error) {

		return req, nil

	})

	assert.Equal(t, res, steams)
	assert.Equal(t, nil, err)
}

func BenchmarkSteams(b *testing.B) {

	ctx := context.TODO()
	p := pool.NewLimited(10)

	defer p.Close()

	for i := 0; i < b.N; i++ {

		_, _ = Steams[int64, int64](ctx, p, steams, func(ctx context.Context, req int64) (interface{}, error) {

			return req, nil

		})

	}
}

func TestTasks(t *testing.T) {

	ctx := context.TODO()
	p := pool.NewLimited(10)
	defer p.Close()

	res, err := Tasks[int64, int64, int64](ctx, p, tasks, func(ctx context.Context, req int64) (interface{}, error) {

		return req, nil

	})

	assert.Equal(t, res, tasks)
	assert.Equal(t, nil, err)
}

func BenchmarkTasks(b *testing.B) {

	ctx := context.TODO()
	p := pool.NewLimited(10)

	defer p.Close()

	for i := 0; i < b.N; i++ {

		_, _ = Tasks[int64, int64, int64](ctx, p, tasks, func(ctx context.Context, req int64) (interface{}, error) {

			return req, nil

		})

	}
}
