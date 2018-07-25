package main

import (
	"fmt"
)

type calculator struct {
	acc float64
}

type opfunc func(float64) float64

func (c *calculator) do(op opfunc) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func add(n float64) opfunc {
	return func(acc float64) float64 {
		return acc + n
	}
}

func sub(n float64) opfunc {
	return func(acc float64) float64 {
		return acc - n
	}
}

func mul(n float64) opfunc {
	return func(acc float64) float64 {
		return acc * n
	}
}

func div(n float64) opfunc {
	return func(acc float64) float64 {
		return acc / n
	}
}

func main() {
	var c calculator
	fmt.Println(c.do(add(10)))
	fmt.Println(c.do(add(30)))
	fmt.Println(c.do(sub(10)))
	fmt.Println(c.do(mul(10)))
	fmt.Println(c.do(div(2)))
}
