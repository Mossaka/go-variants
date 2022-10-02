package examples

//
// variant c1 {
//   a(s32)
//   b(s64)
// }

type C1 interface {
	IsC1()
}

type C1A struct {
	Val int32
}

func (C1A) IsC1() {}

type C1B struct {
	Val int64
}

func (C1B) IsC1() {}

func VariantC1(c1 C1) {
	switch c := c1.(type) {
	case C1A:
		println(c.Val)
	case C1B:
		println(c.Val)
	}
}
