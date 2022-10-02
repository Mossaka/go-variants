package examples

// record r {
//   a: u8,
// }
//
// variant c2 {
//   a(s32)
//   b(s64)
// }
//
// variant nested {
//   a(r)
//   b(c2)
// }

type R struct {
	A uint8
}

type C2 interface {
	IsC2()
}

type C2A struct {
	Val int32
}

func (C2A) IsC2() {}

type C2B struct {
	Val int64
}

func (C2B) IsC2() {}

type Nested interface {
	IsNested()
}

type NestedA struct {
	Val R
}

func (NestedA) IsNested() {}

type NestedB struct {
	Val C2
}

func (NestedB) IsNested() {}

func VariantNested(nested Nested) {
	switch n := nested.(type) {
	case NestedA:
		println(n.Val.A)
	case NestedB:
		switch c := n.Val.(type) {
		case C2A:
			println(c.Val)
		case C2B:
			println(c.Val)
		}
	}
}
