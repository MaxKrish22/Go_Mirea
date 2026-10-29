package main

import "fmt"

func main() {
	c1 := 20
	c2 := 25
	c3 := 35

	var sr_znach float64 = float64(c1+c2+c3) / float64(3)
	fmt.Println("Сумма:", c1+c2+c3)
	fmt.Printf("Среднее : %.2f\n", sr_znach)
}
