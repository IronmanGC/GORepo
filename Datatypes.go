// Basic Data Types
// 1. Integer
// 2. String
// 3. Boolean
package main

import "fmt"

func main() {

	var ia int=10
	ib:=20

	fmt.Println(" :: Integer Type :: ")
	fmt.Println(ia)
	fmt.Println(ib)


	var ss1 string="Welcome in Go"
	ss2:="Have nice day"

	fmt.Println(" :: String Type :: ")
	fmt.Println(ss1)
	fmt.Println(ss2)


	var bi bool = true
	bj:=false

	fmt.Println(" :: Boolean Type :: ")
	fmt.Println(bi)
	fmt.Println(bj)

	fmt.Printf("bi Type %T ",bi)


}