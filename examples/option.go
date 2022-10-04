package examples

// roundtrip-option: func(a: option<u32>) -> option<u8>

type Option[V any] interface {
	IsOption()
}

type OptionSome[V any] struct {
	Val V
}

func (OptionSome[V]) IsOption() {}

type OptionNone struct {
}

func (OptionNone) IsOption() {}

// roundtrip-option: func(a: option<u32>) -> option<u8>

func RoundtripOption(a Option[uint32]) Option[uint8] {
	switch a := a.(type) {
	case OptionSome[uint32]:
		return OptionSome[uint8]{Val: uint8(a.Val)}
	case OptionNone:
		return a
	}
	panic("unreachable")
}
