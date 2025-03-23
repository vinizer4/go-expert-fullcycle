package math

type math struct {
	a int
	b int
}

func (m math) Add() int {
	return m.a + m.b
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}
