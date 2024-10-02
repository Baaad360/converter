package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var i int
	var s1 string
	fmt.Println("masukan value int")
	fmt.Scan(&i)
	fmt.Println("masukkan value string")
	fmt.Scan(&s1)
	s := strconv.Itoa(i) //ini converter int ke string
	i1,_ := strconv.Atoi(s1) //ini converter string ke int

	fmt.Println()
	fmt.Println("ini convert int -> string =",s)
	fmt.Println("ini convert string -> int =",i1)
}
