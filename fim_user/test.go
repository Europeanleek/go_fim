package main

import (
	"fmt"
	"slices"
)

func main() {
	temp_str := "apple"
	temp := []byte(temp_str)
	slices.Sort(temp)
	ortedS := string(temp)
	fmt.Println(ortedS)
}
