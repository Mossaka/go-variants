package examples

// roundtrip-result: func(a: result<u32, u32>) -> result<u64, u8>

type ResultU64U8 interface {
	IsResultU64U8()
}

type ResultU64U8Ok struct {
	Val uint64
}

func (ResultU64U8Ok) IsResultU64U8() {}

type ResultU64U8Err struct {
	Val uint8
}

func (ResultU64U8Err) IsResultU64U8() {}

type ResultU32U32 interface {
	IsResultU32U32()
}

type ResultU32U32Ok struct {
	Val uint32
}

func (ResultU32U32Ok) IsResultU32U32() {}

type ResultU32U32Err struct {
	Val uint32
}

func (ResultU32U32Err) IsResultU32U32() {}

func RoundtripResult(a ResultU32U32) ResultU64U8 {
	switch a := a.(type) {
	case ResultU32U32Ok:
		return ResultU64U8Ok{Val: uint64(a.Val)}
	case ResultU32U32Err:
		return ResultU64U8Err{Val: uint8(a.Val)}
	}
	panic("unreachable")
}
