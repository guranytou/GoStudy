package main

import (
	"fmt"
)

func main() {
	var W, H, x, y, r int
	fmt.Scanf("%d %d %d %d %d", &W, &H, &x, &y, &r)

	if x-r >= 0 && y-r >= 0 && W >= x+r && H >= y+r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
