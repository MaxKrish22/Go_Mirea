package main

import "fmt"

func main() {
	var mins int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&mins)
	if err != nil {
		fmt.Print("Входные данные должны быть числом")
	} else if mins < 0 {
		fmt.Print("Число минут должно быть больше или равно 0")
	} else {
		var hours int = mins / 60
		var minutes int = mins % 60

		fmt.Printf("Полных часов: %d\n", hours)
		fmt.Printf("Оставшихся минут: %d\n", minutes)
	}
}
