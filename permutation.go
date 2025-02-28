package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := n; i > 0; i-- {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(X,Y int) int{
	if X >= Y {
		return faktorial(X) / faktorial(X-Y)
	}else {
		return faktorial(Y) / faktorial(Y-X)

	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Printf("%d %d ",faktorial(x), faktorial(y))
	if x <= y {
		fmt.Println(permutasi(y,x))
	} else if x >= y {
		fmt.Println(permutasi(x,y))
	}
}
