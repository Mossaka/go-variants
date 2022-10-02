package main

import "github.com/mossaka/tinygo-variant/examples"

func main() {
	examples.VariantC1(examples.C1A{Val: 1})
	examples.VariantNested(examples.NestedA{Val: examples.R{A: 1}})
	examples.VariantNested(examples.NestedB{Val: examples.C2A{Val: 1}})
}
