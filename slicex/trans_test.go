package slicex

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type TransStruct struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

type TransStruct2 struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

func (t *TransStruct) ToString() string {

	b, _ := json.Marshal(t)

	return string(b)
}

// @test TestSTransStruct Slice string to Struct
func TestSTransStruct(t *testing.T) {
	var tests []string

	m := []TransStruct{
		{
			Id:       1,
			Name:     "test1",
			CreateAt: time.Now(),
		},
		{
			Id:       2,
			Name:     "test2",
			CreateAt: time.Now(),
		},
	}

	for _, v := range m {

		tests = append(tests, v.ToString())

	}

	transData, _ := STrans[TransStruct](tests)

	for i, v := range transData {

		assert.Equal(t, m[i].Id, v.Id)
		assert.Equal(t, m[i].Name, v.Name)
		assert.Equal(t, m[i].CreateAt.Equal(v.CreateAt), true)

	}

}

// @var test TestSTransInt Slice string to Int
func TestSTransInt(t *testing.T) {

	var (
		testsString = []string{"1", "2", "3", "4", "5"}
		testsInt    = []int{1, 2, 3, 4, 5}
	)

	transData, _ := STrans[int](testsString)

	for i, v := range transData {

		assert.Equal(t, testsInt[i], v)

	}

}

func BenchmarkSTransStruct(b *testing.B) {

	m := []string{
		(&TransStruct{Id: 1, Name: "test1", CreateAt: time.Now()}).ToString(),
		(&TransStruct{Id: 2, Name: "test2", CreateAt: time.Now()}).ToString(),
	}

	for i := 0; i < b.N; i++ {

		_, _ = STrans[TransStruct](m)

	}

}

func BenchmarkJsonUnmarshalStruct(b *testing.B) {

	m := []string{
		(&TransStruct{Id: 1, Name: "test1", CreateAt: time.Now()}).ToString(),
		(&TransStruct{Id: 2, Name: "test2", CreateAt: time.Now()}).ToString(),
	}

	for i := 0; i < b.N; i++ {

		for _, v := range m {

			var (
				tmp TransStruct
			)

			_ = json.Unmarshal([]byte(v), &tmp)

		}

	}

}

func BenchmarkStructToStruct(b *testing.B) {

	a1, b1 := TransStruct{
		Id:   1,
		Name: "test1",
	}, new(TransStruct2)

	for i := 0; i < b.N; i++ {

		b1.Id = a1.Id
		b1.Name = a1.Name

	}

}

func BenchmarkCopierStruct(b *testing.B) {

	a1, b1 := TransStruct{
		Id:   1,
		Name: "test1",
	}, new(TransStruct2)

	for i := 0; i < b.N; i++ {
		copier.Copy(b1, a1)
	}

}
