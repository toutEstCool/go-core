package main

import "fmt"


func main() {
	fmt.Println("Hello World...")
  // Типы данных 
  var a int8 = -1
  var b int16 = -2
  var c int32 = 30
	var d int64 = -9223372036854
	var e uint8 = 230
	var f uint16 = 21000
	var g uint32 = 429496
	var h uint64 = 184467440737095516
	var i byte = 140
	var j rune = 214748
	var k int = 54
	var l uint = 120
	var a1 float32 = 31.1123908
	var a2 float64 = -67123.124216236
	var c1 complex64 = 1+2i
	var c2 complex128 = 4+90i

	var b1 bool = true
	var b2 bool = false

	var name string = "Danil"


	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, a1, a2, c1, c2, b1, b2, name)
}
