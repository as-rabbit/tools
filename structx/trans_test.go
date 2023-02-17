package structx

import (
	"encoding/json"
	"fmt"
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

func (t *TransStruct) Bytes() (res []byte, err error) {

	if nil != t {

		return json.Marshal(t)

	}

	return

}

// @test TestBTransStruct Slice string to Struct
func TestBTrans(t *testing.T) {
	var n *TransStruct

	m := TransStruct{
		Id:       1,
		Name:     "test1",
		CreateAt: time.Now(),
	}

	m2, _ := BTrans[TransStruct2](&m)

	assert.Equal(t, m.Id, m2.Id)
	assert.Equal(t, m.Name, m2.Name)
	assert.Equal(t, m.CreateAt.Equal(m2.CreateAt), true)

	m3, _ := BTrans[*TransStruct2](n)

	fmt.Println(m3)

}

func BenchmarkStructToStruct(b *testing.B) {

	a1, b1 := TransStruct{
		Id:       1,
		Name:     "test1",
		CreateAt: time.Now(),
	}, new(TransStruct2)

	for i := 0; i < b.N; i++ {

		b1.Id = a1.Id
		b1.Name = a1.Name
		b1.CreateAt = a1.CreateAt

	}

}

func BenchmarkAnyToStruct(b *testing.B) {

	a1 := TransStruct{
		Id:   1,
		Name: "test1",
	}

	for i := 0; i < b.N; i++ {

		_, _ = BTrans[TransStruct2](&a1)

	}

}

func BenchmarkCopierStruct(b *testing.B) {

	a1, b1 := TransStruct{
		Id:       1,
		Name:     "test1",
		CreateAt: time.Now(),
	}, new(TransStruct2)

	for i := 0; i < b.N; i++ {
		copier.Copy(b1, a1)
	}
}
