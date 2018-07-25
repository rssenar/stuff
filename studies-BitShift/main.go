package main

import "fmt"

func main() {
	var a int8 = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a<<1)
	fmt.Printf("%08b\n", a<<2)
	fmt.Printf("%08b\n", a<<3)
	fmt.Println("")

	var b uint8 = 120
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", b>>1)
	fmt.Printf("%08b\n", b>>2)
	fmt.Println("")

	var c uint32 = 32
	fmt.Printf("%08b\n", c<<20)
	fmt.Println(c << 20)
	fmt.Println("")

	d := 8
	fmt.Printf("%d\n", d>>1)
	fmt.Printf("%d\n", d>>2)
	fmt.Printf("%d\n", d>>3)
	fmt.Println("")

	e := 8
	fmt.Printf("%d\n", e<<1)
	fmt.Printf("%d\n", e<<2)
	fmt.Printf("%d\n", e<<3)
	fmt.Println("")

	f := 8
	fmt.Printf("%d\n", f<<10)
	fmt.Printf("%d\n", f<<20)
	fmt.Printf("%d\n", f<<30)
}
