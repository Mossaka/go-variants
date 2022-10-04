package examples

import "fmt"

// roundtrip-result: func(a: result<u32, u32>) -> result<u64, u8>

type Result[V any, E any] interface {
	IsResult()
}

type ResultOk[V any] struct {
	Val V
}

func (ResultOk[V]) IsResult() {}

type ResultErr[E any] struct {
	Err E
}

func (ResultErr[E]) IsResult() {}

func RoundtripResult(a Result[uint32, uint32]) Result[uint64, uint8] {
	switch a := a.(type) {
	case ResultOk[uint32]:
		fmt.Printf("ok: %d\n", a.Val)
		return ResultOk[uint64]{Val: uint64(a.Val)}
	case ResultErr[uint32]:
		fmt.Printf("err: %d\n", a.Err)
		return ResultErr[uint8]{Err: uint8(a.Err)}
	}
	panic("unreachable")
}
