package list

import (
	"container/list"
)

type timeElem struct {
	time	uint64
	name	string
}

var (
	ls	*list.List
)

func init() {
	ls = list.New()
}