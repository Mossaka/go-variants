package examples

// roundtrip-option: func(a: option<u32>) -> option<u8>

type OptionU8 interface {
	IsOptionU8()
}

type OptionU8Some struct {
	Val uint8
}

func (OptionU8Some) IsOptionU8() {}

type OptionU8None struct {
}

func (OptionU8None) IsOptionU8() {}

type OptionU32 interface {
	IsOptionU32()
}

type OptionU32Some struct {
	Val uint32
}

func (OptionU32Some) IsOptionU32() {}

type OptionU32None struct {
}

func (OptionU32None) IsOptionU32() {}

// roundtrip-option: func(a: option<u32>) -> option<u8>

func RoundtripOption(a OptionU32) OptionU8 {
	switch a := a.(type) {
	case OptionU32Some:
		return OptionU8Some{Val: uint8(a.Val)}
	case OptionU32None:
		return OptionU8None{}
	}
	panic("unreachable")
}
