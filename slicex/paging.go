package slicex

import (
	"golang.org/x/exp/constraints"
	"math"
)

// PaginationResp Return Pagination Resp
type PaginationResp[T constraints.Integer, D comparable] struct {
	IndexStart T
	IndexEnd   T
	PageTotal  T
	Data       []D
}

// Empty Return Data
func emptyPageResp[T constraints.Integer, D comparable]() (res *PaginationResp[T, D]) {
	return &PaginationResp[T, D]{
		IndexStart: 0,
		IndexEnd:   0,
		PageTotal:  1,
		Data:       []D{},
	}
}

func overstepPageResp[T constraints.Integer, D comparable](pt T) (res *PaginationResp[T, D]) {
	return &PaginationResp[T, D]{
		IndexStart: 0,
		IndexEnd:   0,
		PageTotal:  pt,
		Data:       []D{},
	}
}

// Top Return Slice Top
func Top[T constraints.Integer, D comparable](top T, p []D) (res []D) {

	var (
		indexStart T
		indexEnd   T
		total      = T(len(p))
	)

	res = make([]D, top)

	if 0 >= top {

		return p

	}

	indexStart, indexEnd = 0, top

	if indexEnd > total {

		indexEnd = total

	}

	copy(res, p[indexStart:indexEnd])

	return res
}

func Pagination[T constraints.Integer, D comparable](nowPage T, pageSize T, data []D) (res *PaginationResp[T, D]) {
	var (
		pageTotal  T
		indexStart T
		indexEnd   T
		total      = T(len(data))
	)

	if 0 >= nowPage || 0 >= pageSize {

		return emptyPageResp[T, D]()

	}

	if 0 >= total {

		return emptyPageResp[T, D]()

	}

	pageTotal = T(math.Ceil(float64(total) / float64(pageSize)))

	if nowPage > pageTotal {

		return overstepPageResp[T, D](pageTotal)

	}

	indexStart, indexEnd = (nowPage-1)*pageSize, indexStart+pageSize

	if indexEnd > total {
		indexEnd = total
	}

	pageData := data[indexStart:indexEnd]
	resp := make([]D, len(pageData), pageSize)

	copy(resp, pageData)

	return &PaginationResp[T, D]{
		IndexStart: indexStart,
		IndexEnd:   indexEnd,
		PageTotal:  pageTotal,
		Data:       resp,
	}
}
