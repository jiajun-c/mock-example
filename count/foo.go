package count

type Foo struct {
	bar Bar
}

func NewFoo(bar Bar) *Foo {
	return &Foo{bar: bar}
}
