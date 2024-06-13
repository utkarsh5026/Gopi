package main

import (
	"fmt"
	"gopi/array"
)

func main() {
	arr := array.Fill(9, 4, 5, 6)
	fmt.Println(*arr)
}
